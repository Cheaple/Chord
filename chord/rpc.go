package chord

// RPCs for Chord nodes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"net"

	"chord/chord/rpc"
)

func (n *Node) StartRPCService() {
	server := grpc.NewServer()
	rpc.RegisterChordServer(server, n)
	n.rpcService.server = server

	// Start server at the node's address
	listener, err := net.Listen("tcp", string(n.Address))
	if err != nil {
		log.Fatalf("Error listening at %s: %v", n.Address, err)
	}
	fmt.Println("Chord node starts listening at %s", n.Address)
	go server.Serve(listener)
}


//
// Find the successor node of a given ID in the Chord ring
// starting searching from a give address
//
func (n *Node) findSuccessorRPC(ety *NodeEntry, id *big.Int) (*rpc.NodeEntry, error) {
	conn, err := grpc.Dial(string(ety.Address))
	if err != nil {
		return nil, err
	}
	client := rpc.NewChordClient(conn)

	req := &rpc.EmptyMsg{}
	ctx := context.Background()
	return client.GetSuccessor(ctx, req)
}

//
// Find the predecessor node of the given node 
//
func (n *Node) findPredecessorRPC(ety *NodeEntry) (*rpc.NodeEntry, error) {
	conn, err := grpc.Dial(string(ety.Address))
	if err != nil {
		return nil, err
	}
	client := rpc.NewChordClient(conn)

	req := &rpc.EmptyMsg{}
	ctx := context.Background()
	return client.GetPredecessor(ctx, req)
}

/* ******************************************************************************* *
 * ************************* RPC Interface Implementaiton************************* */

func (n *Node) GetPredecessor(ctx context.Context, in *rpc.EmptyMsg) (*rpc.NodeEntry, error) {
	return n.Predecessor.toRPC(), nil
}

func (n *Node) GetSuccessor(ctx context.Context, in *rpc.EmptyMsg) (*rpc.NodeEntry, error) {
	return n.Successors[1].toRPC(), nil
}