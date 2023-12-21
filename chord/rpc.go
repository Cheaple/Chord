package chord

// RPCs for Chord nodes

import (
	"math/big"
	// "google.golang.org/grpc"
)


func call(address string, method string, request interface{}, reply interface{}) error {

	
	return nil
}




//
// Find the successor node of a given ID in the Chord ring
// starting searching from a give address
//
func (n *Node) findSuccessorRPC(ety NodeEntry, id *big.Int) (NodeEntry, error) {
	
	
	return NodeEntry{}, nil
}