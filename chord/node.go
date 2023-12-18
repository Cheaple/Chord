package chord

import (
	"crypto/sha1"

	"chord/utils"
)


const m = 6  // m-bit identifier
var fingerTableSize = m

type Key string

type NodeAddress string

type Node struct {
	Name       	string   			// Name: IP:Port or User specified Name. Exp: [N]14
	Identifier	*big.Int 			// Hash(Address) -> Chord space Identifier

    Address    	NodeAddress
    FingerTable	[]NodeAddress
    Predecessor	NodeAddress
    Successors 	[]NodeAddress

    Bucket map[Key]string

}


type NodeInfoEntry struct {
	Identifier	int
	Address		NodeAddress
}

func hashString(elt string) *big.Int {
    hasher := sha1.New()
    hasher.Write([]byte(elt))
    return new(big.Int).SetBytes(hasher.Sum(nil))
}

// Create a new nod
func NewNode(args utils.Arguments) *Node {
	node := &Node{}

	node.Address = NodeAddress(fmt.Sprintf("%s:%d", args.Address, args.Port))
	if args.Identifier == "" {
		
	} else {
		node.Identifier
	}


	node.FingerTable = make([]NodeInfoEntry, fingerTableSize + 1)
	node.Successors = make([]NodeInfoEntry, args.CntSuccessors)
	
	node.Bucket = make(map[int]string)	

	// TODO: init FingerTable
	// TODO: init Successors
}