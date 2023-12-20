package chord

import (
	"crypto/sha1"
	"math/big"
)


const M = 6  // m-bit identifier

type Key string

type NodeAddress string

type Node struct {
	Name       	string   			// Name: IP:Port or User specified Name. Exp: [N]14
	Identifier	*big.Int 			// Hash(Address) -> Chord space Identifier

    Address    	NodeAddress
    FingerTable	[]NodeEntry
    Predecessor	NodeAddress
    Successors 	[]NodeEntry

    Bucket 		map[Key]string

	doneCh		chan struct{}
}

type NodeEntry struct {
	Id			[]byte
	Address		NodeAddress
}

// func (entry *NodeEntry) Set(id *bigInt, address NodeAddress) {
// 	entry.Id.Set(id)
// 	entry.Address = address
// }

type NodeTable []NodeEntry

func hashString(elt string) *big.Int {
    hasher := sha1.New()
    hasher.Write([]byte(elt))
    return new(big.Int).SetBytes(hasher.Sum(nil))
}