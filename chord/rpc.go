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


//
// Find the successor node of a given ID in the Chord ring
// starting searching from a give address
//
func (n *Node) findSuccessorRPC(ety *NodeEntry, id *big.Int) (*NodeEntry, error) {
	conn, err := grpc.Dial(string(ety.Address))
	if err != nil {
		return nil, err
	}
	client := NewChordClient(conn)

	req := &EmptyMsg{}
	ctx := context.Background()
	return client.GetSuccessor(ctx, req)
}

//
// Find the predecessor node of the given node 
//
func (n *Node) findPredecessorRPC(ety *NodeEntry) (*NodeEntry, error) {
	conn, err := grpc.Dial(string(ety.Address))
	if err != nil {
		return nil, err
	}
	client := NewChordClient(conn)

	req := &EmptyMsg{}
	ctx := context.Background()
	return client.GetPredecessor(ctx, req)
}

/* ******************************************************************************* *
 * ************************* RPC Interface Implementaiton************************* */

func (n *Node) GetPredecessor(ctx context.Context, in *EmptyMsg) (*NodeEntry, error) {
	return n.Predecessor, nil
}

func (n *Node) GetSuccessor(ctx context.Context, in *EmptyMsg) (*NodeEntry, error) {
	return n.Successors[1], nil
}