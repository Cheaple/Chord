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
func findSuccessorRPC(ety *NodeEntry, id *big.Int) (*NodeEntry, error) {
	
	
	return nil, nil
}

func findPredecessorRPC(ety *NodeEntry) (*NodeEntry, error) {


	return nil, nil
}

/* ******************************************************************************* *
 * ************************* RPC Interface Implementaiton************************* */

func (n *Node) GetPredecessor(ctx context.Context, in *rpc.Empty) (*rpc.NodeEntry, error) {


	return nil, nil
}

func (n *Node) GetSuccessor(ctx context.Context, in *rpc.Empty) (*rpc.NodeEntry, error) {


	return nil, nil
}