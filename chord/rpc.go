package chord

// RPCs for Chord nodes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"net"
)

func (n *Node) StartRPCService() {
	server := grpc.NewServer()
	RegisterChordServer(server, n)
	n.rpcService.server = server

	// Start server at the node's address
	listener, err := net.Listen("tcp", string(n.Address))
	if err != nil {
		log.Fatalf("Error listening at %s: %v", n.Address, err)
	}
	fmt.Println("Chord node starts listening at %s", n.Address)
	go server.Serve(listener)
}

/* ******************************************************************************* *
 * *********************************** RPC Calls ********************************* */

//
// Locate the successor node of a given ID in the Chord ring
// starting searching from a give address
//
func (n *Node) LocateRPC(ety *NodeEntry, id *big.Int) (*NodeEntry, error) {
	n.DPrintf("LocateRPC(): target address = %s", string(ety.Address))
	conn, err := grpc.Dial(string(ety.Address), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := NewChordClient(conn)

	req := &BytesMsg{ Data: id.Bytes() }
	ctx := context.Background()
	return client.Locate(ctx, req)
}

//
// Find the successor list of the given node 
//
func (n *Node) GetSuccessorListRPC(ety *NodeEntry) (*NodeList, error) {
	n.DPrintf("GetSuccessorListRPC(): target address = %s", string(ety.Address))
	conn, err := grpc.Dial(string(ety.Address), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := NewChordClient(conn)

	req := &EmptyMsg{}
	ctx := context.Background()
	return client.GetSuccessorList(ctx, req)
}

//
// Find the predecessor node of the given node 
//
func (n *Node) GetPredecessorRPC(ety *NodeEntry) (*NodeEntry, error) {
	n.DPrintf("GetPredecessorRPC(): target node = %s", ety.Address)
	conn, err := grpc.Dial(string(ety.Address), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := NewChordClient(conn)

	req := &EmptyMsg{}
	ctx := context.Background()
	return client.GetPredecessor(ctx, req)
}

//
// Notify the given node to set a new predecessor
//
func (n *Node) NotifyRPC(ety *NodeEntry) (*EmptyMsg, error) {
	n.DPrintf("NotifyRPC(): target node = %s", ety.Address)
	conn, err := grpc.Dial(string(ety.Address), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := NewChordClient(conn)

	req := n.Entry
	ctx := context.Background()
	return client.SetPredecessor(ctx, req)
}

//
// Check whether the predecessor fails
//
func (n *Node) CheckRPC(ety *NodeEntry) (*EmptyMsg, error) {
	n.DPrintf("CheckRPC(): target node = %s", ety.Address)
	conn, err := grpc.Dial(string(ety.Address), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := NewChordClient(conn)

	req := &EmptyMsg{}
	ctx := context.Background()
	return client.Check(ctx, req)
}

//
// Check whether a file exists
//
func (n *Node) CheckKeyRPC(ety *NodeEntry, key string) (bool, error) {
	n.DPrintf("CheckKeyRPC(): target node = %s, key = %s", ety.Address, key)
	conn, err := grpc.Dial(string(ety.Address), grpc.WithInsecure())
	if err != nil {
		return false, err
	}
	client := NewChordClient(conn)

	req := &KeyMsg{Key: key,}
	ctx := context.Background()
	boolMsg, err := client.CheckKey(ctx, req)
	return boolMsg.Success, err
}

/* ******************************************************************************* *
 * ******************************* RPC Responses ********************************* */

// When receiving RPC calls, nodes run the following functions to generate RPC responses

func (n *Node) Locate(ctx context.Context, in *BytesMsg) (*NodeEntry, error) {
	n.DPrintf("Check()")
	return n.locateSuccessor(new(big.Int).SetBytes(in.Data))
} 

func (n *Node) Check(ctx context.Context, in *EmptyMsg) (*EmptyMsg, error) {
	n.DPrintf("Check()")
	return &EmptyMsg{}, nil
} 

func (n *Node) GetPredecessor(ctx context.Context, in *EmptyMsg) (*NodeEntry, error) {
	n.DPrintf("GetPredecessor()")
	return n.Predecessor, nil
}

func (n *Node) GetSuccessorList(ctx context.Context, in *EmptyMsg) (*NodeList, error) {
	n.DPrintf("GetSuccessorList()")
	return n.Successors, nil
}

func (n *Node) SetPredecessor(ctx context.Context, pred *NodeEntry) (*EmptyMsg, error) {
	n.DPrintf("SetPredecessor(): %+v", pred)
	if n.Predecessor.empty() || nodeBetweenOpen(n.Predecessor, pred, n.Entry) {
		n.DPrintf("SetPredecessor(): set predecessor = %s", pred.Address)
		n.Predecessor = pred
	}
	return &EmptyMsg{}, nil
}

func (n *Node) CheckKey(ctx context.Context, in *KeyMsg) (*BoolMsg, error) {
	n.DPrintf("CheckKey(): %+v", in)
	_, ok := n.Bucket[in.Key]
	return &BoolMsg{ Success: ok }, nil
}