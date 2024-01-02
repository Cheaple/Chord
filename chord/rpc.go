package chord

// RPCs for Chord nodes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"net"
)

const chunkSize = 4096  // chunk size when transferring files through RPCs

func (n *Node) startRPCService() {
	// Start server
	server := grpc.NewServer()
	RegisterChordServer(server, n)
	n.rpcService.server = server

	// Server starts listening at the node's address
	listener, err := net.Listen("tcp", string(n.Address))
	if err != nil {
		log.Fatalf("Error listening at %s: %v", n.Address, err)
	}
	fmt.Println("Chord node starts listening at %s", n.Address)
	go server.Serve(listener)
}

//
// Build a TCP connection for normal RPCs
//
func (n *Node) makeClient(ety *NodeEntry) (ChordClient, error) {
	conn, err := grpc.Dial(string(ety.Address), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return NewChordClient(conn), nil
}


/* ******************************************************************************* *
 * *********************************** RPC Calls ********************************* */

//
// Locate the successor node of a given ID in the Chord ring
// starting searching from a give address
//
func (n *Node) LocateRPC(ety *NodeEntry, id *big.Int) (*NodeEntry, error) {
	n.DPrintf("LocateRPC(): target address = %s", string(ety.Address))

	// Build a connection
	client, err := n.makeClient(ety)
	if err != nil {
		return nil, err
	}

	req := &BytesMsg{ Data: id.Bytes() }
	ctx := context.Background()
	return client.Locate(ctx, req)
}

//
// Find the successor list of the given node 
//
func (n *Node) GetSuccessorListRPC(ety *NodeEntry) (*NodeList, error) {
	n.DPrintf("GetSuccessorListRPC(): target address = %s", string(ety.Address))

	// Build a connection
	client, err := n.makeClient(ety)
	if err != nil {
		return nil, err
	}

	req := &EmptyMsg{}
	ctx := context.Background()
	return client.GetSuccessorList(ctx, req)
}

//
// Find the predecessor node of the given node 
//
func (n *Node) GetPredecessorRPC(ety *NodeEntry) (*NodeEntry, error) {
	n.DPrintf("GetPredecessorRPC(): target node = %s", ety.Address)

	// Build a connection
	client, err := n.makeClient(ety)
	if err != nil {
		return nil, err
	}

	req := &EmptyMsg{}
	ctx := context.Background()
	return client.GetPredecessor(ctx, req)
}

//
// Notify the given node to set a new predecessor
//
func (n *Node) NotifyRPC(ety *NodeEntry) (bool, error) {
	n.DPrintf("NotifyRPC(): target node = %s", ety.Address)

	// Build a connection
	client, err := n.makeClient(ety)
	if err != nil {
		return false, err
	}

	req := n.Entry
	ctx := context.Background()
	boolMsg, err := client.SetPredecessor(ctx, req)
	return boolMsg.Success, err
}

//
// Check whether the predecessor fails
//
func (n *Node) CheckRPC(ety *NodeEntry) (bool, error) {
	n.DPrintf("CheckRPC(): target node = %s", ety.Address)

	// Build a connection
	client, err := n.makeClient(ety)
	if err != nil {
		return false, err
	}

	req := &EmptyMsg{}
	ctx := context.Background()
	boolMsg, err := client.Check(ctx, req)
	if err != nil {
		return false, err
	}
	return boolMsg.Success, err
}

//
// Check whether a file exists
//
func (n *Node) CheckKeyRPC(ety *NodeEntry, key string) (bool, error) {
	n.DPrintf("CheckKeyRPC(): target node = %s, key = %s", ety.Address, key)

	// Build a connection
	client, err := n.makeClient(ety)
	if err != nil {
		return false, err
	}

	req := &StringMsg{Str: key}
	ctx := context.Background()
	boolMsg, err := client.CheckKey(ctx, req)
	if err != nil {
		return false, err
	}
	return boolMsg.Success, err
}

//
// Store a file in the target node
//
func (n *Node) UploadFileRPC(ety *NodeEntry, filePath string) (bool, error) {
	n.DPrintf("UploadFileRPC(): target node = %s, filePath = %s", ety.Address, filePath)

	// Build a secure connection
	client, err := n.makeClient(ety)
	if err != nil {
		return false, err
	}

	// Open the local file
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("Error opening file %s: %v", filePath, err)
	}
	defer file.Close()

	// Open a stream-based connection 
	ctx := context.Background()
	stream, err := client.UploadFile(ctx)
	
	// Allocate a buffer with `chunkSize` as the capacity
	buffer := make([]byte, chunkSize)

	// Send file data by chunk
	for {
		// Read a chunk from the local file
		bytesRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, fmt.Errorf("Error reading file %s: %v", filePath, err)
		}

		// Send a chunk to the target node
		if err := stream.Send(&FileMsg{
			Name:    filepath.Base(filePath),
			Content: buffer[:bytesRead],
		}); err != nil {
			return false, fmt.Errorf("Error sending file %s: %v", filePath, err)
		}
	}

	// Receive a response from the target node
	boolMsg, err := stream.CloseAndRecv()
	if err != nil {
		return false, err
	}
	if boolMsg.Success == false {
		return boolMsg.Success, fmt.Errorf(boolMsg.ErrorMsg)
	}
	return true, nil
}

//
// Download a file from the target node
//
func (n *Node) DownloadFileRPC(ety *NodeEntry, fileName string) (bool, error) {
	n.DPrintf("UploadFileRPC(): target node = %s, filePath = %s", ety.Address, fileName)

	// Build a connection
	client, err := n.makeClient(ety)
	if err != nil {
		return false, err
	}

	// Open a stream-based connection 
	req := &StringMsg{Str: fileName}
	ctx := context.Background()
	stream, err := client.DownloadFile(ctx, req)
	
	// Create a local temp file
	n.DPrintf("create a new file: %s\n", fileName)
	tmpFile, err := ioutil.TempFile(".", "tmp-" + fileName)
	if err != nil {
		return false, fmt.Errorf("Error creating temp file:", err)
	}

	// Receive uploaded file chunk by chunk
	for idx := 0; ; idx++ {
		// Receive one chunk
		fileRequest, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Remove(tmpFile.Name())
			return false, fmt.Errorf("Error receiving file chunk:", err)
		}
		
		// Write a chunk to the temp file
		n.DPrintf("%s receiving the %d chunk", tmpFile.Name(), idx)
		_, err = tmpFile.Write(fileRequest.Content)
		if err != nil {
			os.Remove(tmpFile.Name())
			return false, fmt.Errorf("Error writing to tmp file: %v", err)
		}
	}
	
	// Save temp file to the local data store
	err = os.Rename(tmpFile.Name(), fileName)
	n.DPrintf("rename temp file '%s' to local file '%s'", tmpFile.Name, fileName)
	if err != nil {
		os.Remove(tmpFile.Name())
		return false, fmt.Errorf("Error saving tmp file: %v", err)
	}
	return true, nil
}


/* ******************************************************************************* *
 * ******************************** RPC Responses ******************************** */

// When receiving RPC calls, nodes run the following functions to generate RPC responses

//
// Handler for Locate rpc service
// Reply location of target identifier in the Chord ring
//
func (n *Node) Locate(ctx context.Context, in *BytesMsg) (*NodeEntry, error) {
	n.DPrintf("Locate()")
	return n.locateSuccessor(new(big.Int).SetBytes(in.Data))
} 

// 
// Handler for Check rpc service
// Reply that this node is alive
//
func (n *Node) Check(ctx context.Context, in *EmptyMsg) (*BoolMsg, error) {
	n.DPrintf("Check()")
	return &BoolMsg{ Success: true }, nil
}

//
// Handler for GetPredecessor rpc service
// Reply this node's predecessor
//
func (n *Node) GetPredecessor(ctx context.Context, in *EmptyMsg) (*NodeEntry, error) {
	n.DPrintf("GetPredecessor()")
	return n.Predecessor, nil
}

//
// Handler for GetSuccessorList rpc service
// Reply this node's successor list
//
func (n *Node) GetSuccessorList(ctx context.Context, in *EmptyMsg) (*NodeList, error) {
	n.DPrintf("GetSuccessorList()")
	return n.Successors, nil
}

//
// Handler for SetPredecessor rpc service
// Set this node's predecessor
//
func (n *Node) SetPredecessor(ctx context.Context, pred *NodeEntry) (*BoolMsg, error) {
	n.DPrintf("SetPredecessor(): %+v", pred)
	if n.Predecessor.empty() || nodeBetweenOpen(n.Predecessor, pred, n.Entry) {
		n.DPrintf("SetPredecessor(): set predecessor = %s", pred.Address)
		n.Predecessor = pred
		return &BoolMsg{ Success: true }, nil
	}
	return &BoolMsg{ Success: false }, nil
}

// Reply whether the key exists in the local data store
func (n *Node) CheckKey(ctx context.Context, in *StringMsg) (*BoolMsg, error) {
	n.DPrintf("CheckKey(): %+v", in)
	_, ok := n.Bucket[in.Str]
	return &BoolMsg{ Success: ok }, nil
}

//
// Handler for UploadFile rpc service
// Store the uploaded file
//
func (n *Node) UploadFile(stream Chord_UploadFileServer) error {
	n.DPrintf("UploadFile() to handle UploadFileRPC")
	filePath := ""
	fileName := ""
	var tmpFile *os.File
	defer tmpFile.Close()

	// Receive uploaded file chunk by chunk
	for idx := 0; ; idx++ {
		// Receive one chunk
		fileRequest, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			stream.SendAndClose(&BoolMsg{Success: false})
			os.Remove(tmpFile.Name())
			return fmt.Errorf("Error receiving file chunk:", err)
		}

		// If first chunk, create a local temp file
		if filePath == "" {
			fileName = fileRequest.Name
			filePath = n.getFilePath(fileName)

			// Check if the file already exists in the node's data store
			_, ok := n.Bucket[fileName]
			if ok {
				return stream.SendAndClose(
					&BoolMsg{
						Success: false,
						ErrorMsg: fmt.Sprintf("file already exists."),
					},
				)
			}

			// Create a local temp file
			n.DPrintf("create a new file: %s\n", fileName)
			tmpFile, err = ioutil.TempFile(".", "tmp-" + fileName)
			if err != nil {
				stream.SendAndClose(&BoolMsg{Success: false})
				return fmt.Errorf("Error creating temp file:", err)
			}
		}
		
		// Write a chunk to the temp file
		n.DPrintf("%s receiving the %d chunk", tmpFile.Name(), idx)
		_, err = tmpFile.Write(fileRequest.Content)
		if err != nil {
			os.Remove(tmpFile.Name())
			stream.SendAndClose(&BoolMsg{Success: false})
			return fmt.Errorf("Error writing to tmp file: %v", err)
		}
	}
	
	// Save temp file to the local data store
	err := os.Rename(tmpFile.Name(), filePath)
	n.DPrintf("rename temp file '%s' to local file '%s'", tmpFile.Name, filePath)
	if err != nil {
		return fmt.Errorf("Error saving tmp file: %v", err)
	}
	n.Bucket[fileName] = 1
	return stream.SendAndClose(&BoolMsg{Success: true})
}

//
// Handler for DownloadFile rpc service
// Transfer the asked file
//
func (n *Node) DownloadFile(in *StringMsg, stream Chord_DownloadFileServer) error {
	n.DPrintf("DownloadFile(): %+v", in)
	fileName := in.Str
	filePath := n.getFilePath(fileName)

	// Check if the file exists in the node's data store
	_, ok := n.Bucket[fileName]
	if ok == false {
		return fmt.Errorf("File not found")
	}

	// Open the local file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error opening file %s: %v", filePath, err)
	}
	defer file.Close()
	
	// Allocate a buffer with `chunkSize` as the capacity
	buffer := make([]byte, chunkSize)

	// Send file data by chunk
	for {
		// Read a chunk from the local file
		bytesRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Error reading file %s: %v", filePath, err)
		}

		// Send a chunk to the target node
		if err := stream.Send(&FileMsg{
			Name:    fileName,
			Content: buffer[:bytesRead],
		}); err != nil {
			return fmt.Errorf("Error sending file %s: %v", filePath, err)
		}
	}

	return nil
}