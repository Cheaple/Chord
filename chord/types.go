package chord

import (
	"crypto/sha1"
    "google.golang.org/grpc"
	"math/big"
)


const M = 7  // m-bit identifier

type Key string

type NodeAddress string

type Node struct {
	// Name       	string   		// Name: IP:Port or User specified Name. Exp: [N]14
	Id      	*big.Int 			// Hash(Address) -> Chord space Identifier

    Address    	NodeAddress         // local address
    FingerTable	[]NodeEntry         // Finger Table
    Predecessor	*NodeEntry          // Predecessor
    Successors 	[]NodeEntry         // Successor List

    Bucket 		map[Key]string      // Buckets to store files

    rpcService  GRPCService         // Service for communications between Chord nodes

	doneCh		chan struct{}       // channel to notify sub-routines to shutdown
    verbose     bool                // whether to print log
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

type GRPCService struct {
    server *grpc.Server
}   

/* ******************************************************************************* *
 * ********************************* Type Operations ***************************** */
 
func hashString(elt string) *big.Int {
    hasher := sha1.New()
    hasher.Write([]byte(elt))
    hash := new(big.Int).SetBytes(hasher.Sum(nil))
    hashMod := new(big.Int).Exp(big.NewInt(2), big.NewInt(M), nil)
    return new(big.Int).Mod(hash, hashMod)
}

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