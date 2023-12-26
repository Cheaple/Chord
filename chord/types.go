package chord

import (
	"crypto/sha1"
    "google.golang.org/grpc"
	"math/big"
)


const M = 7  // m-bit identifier
var hashMod = new(big.Int).Exp(big.NewInt(2), big.NewInt(M), nil)

type Key string

type NodeAddress string

type Node struct {
	// Name       	string   		// Name: IP:Port or User specified Name. Exp: [N]14
	Id      	*big.Int 			// Hash(Address) -> Chord space Identifier
    Address    	NodeAddress         // local address
    Entry       *NodeEntry          // Node Entry

    FingerTable	[]*NodeEntry        // Finger Table
    Predecessor	*NodeEntry          // Predecessor
    Successors 	[]*NodeEntry        // Successor List

    Bucket 		map[Key]string      // Buckets to store files

    rpcService  GRPCService         // Service for communications between Chord nodes

	doneCh		chan struct{}       // channel to notify sub-routines to shutdown
    verbose     bool                // whether to print log
}

// Use NodeEntry struct generated by protoc-buffer, so comment the following definition
//
// type NodeEntry struct {
// 	Identifier	[]byte
// 	Address		NodeAddress
// }

// func (entry *NodeEntry) Set(id *bigInt, address NodeAddress) {
// 	entry.Id.Set(id)
// 	entry.Address = address
// }

type NodeTable []*NodeEntry

type GRPCService struct {
    server *grpc.Server
}   

/* ******************************************************************************* *
 * ********************************* Type Operations ***************************** */

//
// Init NodeEntry
//
func newNodeEntry(id *big.Int, address NodeAddress) *NodeEntry {
	entry := &NodeEntry{}
    entry.Identifier = id.Bytes()
    entry.Address = string(address)
    return entry
}

// Return true if the NodeEntry is empty
func (entry *NodeEntry) empty() bool {
    return entry.Address == ""
}

//
// Init Finger Table or Successor List
//
func (n *Node) newNodeTable(size int) NodeTable {
	tbl := make([]*NodeEntry, size)
	for i := range tbl {
		tbl[i] = n.Entry
        // tbl[i] = &NodeEntry{}
	}
	return tbl
}

// Calculate a given string's hash value
func hashString(elt string) *big.Int {
    hasher := sha1.New()
    hasher.Write([]byte(elt))
    hash := new(big.Int).SetBytes(hasher.Sum(nil))
    return new(big.Int).Mod(hash, hashMod)
}

// Calculate a given NodeAddress's hash value
func hashAddress(address NodeAddress) *big.Int {
    return hashString(string(address))
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

//
// Return true if elt in (start, end)
//
func nodeBetweenOpen(start, elt, end *NodeEntry) bool {
    if elt.empty() {
        return false
    }

    left := new(big.Int).SetBytes(start.Identifier)
    mid := new(big.Int).SetBytes(elt.Identifier)
    right := new(big.Int).SetBytes(end.Identifier)
    return between(left, mid, right, false)
}

//
// Return true if elt in (start, end]
//
func nodeBetweenClosed(start, elt, end *NodeEntry) bool {
    if elt.empty() {
        return false
    }
    
    left := new(big.Int).SetBytes(start.Identifier)
    mid := new(big.Int).SetBytes(elt.Identifier)
    right := new(big.Int).SetBytes(end.Identifier)
    return between(left, mid, right, true)
}

