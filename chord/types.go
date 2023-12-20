package chord

import (
	"crypto/sha1"
	"math/big"
)


const M = 6  // m-bit identifier

type Key string

type NodeAddress string

type Node struct {
	// Name       	string   			// Name: IP:Port or User specified Name. Exp: [N]14
	Id      	*big.Int 			    // Hash(Address) -> Chord space Identifier

    Address    	NodeAddress
    FingerTable	[]NodeEntry
    Predecessor	NodeAddress
    Successors 	[]NodeEntry

    Bucket 		map[Key]string

	doneCh		chan struct{}
}

type NodeEntry struct {
	Identifier	[]byte
	Address		NodeAddress
}

// func (entry *NodeEntry) Set(id *bigInt, address NodeAddress) {
// 	entry.Id.Set(id)
// 	entry.Address = address
// }

type NodeTable []NodeEntry

/* ******************************************************************************* *
 * ********************************* Type Operations ***************************** */
 
func hashString(elt string) *big.Int {
    hasher := sha1.New()
    hasher.Write([]byte(elt))
    return new(big.Int).SetBytes(hasher.Sum(nil))
}

//
// Returns true if elt is between start and end on the ring, 
// accounting for the boundary where the ring loops back on itself. 
// If inclusive is true, it tests if elt is in (start,end], otherwise it tests for (start,end)
//
func between(start, elt, end *big.Int, inclusive bool) bool {
    if end.Cmp(start) > 0 {
        return (start.Cmp(elt) < 0 && elt.Cmp(end) < 0) || (inclusive && elt.Cmp(end) == 0)
    } else {
        return start.Cmp(elt) < 0 || elt.Cmp(end) < 0 || (inclusive && elt.Cmp(end) == 0)
    }
}