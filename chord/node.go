package chord

import (
	"crypto/sha1"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

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
    FingerTable	[]NodeEntry
    Predecessor	NodeAddress
    Successors 	[]NodeEntry

    Bucket map[Key]string


}


type NodeEntry struct {
	Identifier	int
	Address		NodeAddress
}

func hashString(elt string) *big.Int {
    hasher := sha1.New()
    hasher.Write([]byte(elt))
    return new(big.Int).SetBytes(hasher.Sum(nil))
}

//
// Create a new Chord n,
// and create or join a Chor ring
//
func MakeNode(args utils.Arguments) *Node {
	node := &Node{}

	node.Address = NodeAddress(fmt.Sprintf("%s:%d", args.Address, args.Port))
	if args.Identifier == "" {
		node.Identifier = hashString(string(node.Address))
	} else {
		node.Identifier = new(big.Int)
		node.Identifier.SetString(args.Identifier, 16)  // string of 16-bit int to big.Int
	}


	node.FingerTable = make([]NodeEntry, fingerTableSize + 1)
	node.Successors = make([]NodeEntry, args.CntSuccessors)
	node.Bucket = make(map[Key]string)	

	// TODO: init FingerTable
	// TODO: init Successors

	// Join or Create a Chord ring
	if args.JoinAddress != "" {
		joinAddress := NodeAddress(fmt.Sprintf("%s:%d", args.JoinAddress, args.JoinPort))
		fmt.Println("Joining a Chord ring at node", joinAddress)
		err := node.join(joinAddress)
		if err != nil {
			fmt.Println("Error joining the Chord ring:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Creating a new Chord ring at node", node.Address)
		node.create()
	}

	// Peridoically stabilize
	go func() {
		ticker := time.NewTicker(time.Duration(args.StabilizeTime) * time.Millisecond)
		for {
			select {
			case <- ticker.C:
				node.stabilize()
			case <- node.doneCh:
				ticker.Stop()
				return
			}
		}
	}()

	// Peridoically fix finger tables
	go func() {
		next := 0
		ticker := time.NewTicker(time.Duration(args.FixFingerTime) * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				next = node.fixFinger(next)
			case <-node.doneCh:
				ticker.Stop()
				return
			}
		}
	}()

	// Peridoically checkes predecessor	
	go func() {
		ticker := time.NewTicker(time.Duration(args.CheckPredTime) * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				node.checkPredecessor()
			case <-node.doneCh:
				ticker.Stop()
				return
			}
		}
	}()

	return node
}

//
// Create a new Chord ring
//
func (n *Node) create() {
	n.Predecessor = ""
	// n.Predecessor = n.Address

	// init successor list
	for i := 0; i < len(n.Successors); i++ {
		n.Successors[i].Id = n.Identifier.Bytes()
		n.Successors[i].Addr = n.Address
	}
}

//
// Join an old Chord ring
//
func (n *Node) join(joined NodeAddress) error {
	n.Predecessor = ""


}

//
// Print current state
//
func (n *Node) print() {
	fmt.Println("Node Address:", n.Address)
	fmt.Println("Node Identifier:", new(big.Int).SetBytes(n.Identifier.Bytes()))
	fmt.Println("Node Predecessor:", node.Predecessor)

	fmt.Println("----- Successor List -----")
	fmt.Println("Successors  |  Identifier  |  Address ")
	for i := 0; i < len(n.Successors); i++ {
		enrty := n.Successors[i]
		id := new(big.Int).SetBytes(enrty.Id)
		address := enrty.Addr
		fmt.Printf("%6d  |  %10d  |  %s\n", i, id, address)
	}
	
	fmt.Println("----- Finger Table ----")
	fmt.Println("Finger  |  Identifier  |  Address ")
	for i := 0; i < len(n.FingerTable); i++ {
		enrty := n.FingerTable[i]
		id := new(big.Int).SetBytes(enrty.Id)
		address := enrty.Addr
		fmt.Printf("%6d  |  %10d  |  %s\n", i, id, address)
	}

	fmt.Println("----- Buckets -----")
	fmt.Println("Key  |  Identifier  |  Address ")
	for k, v := range node.Bucket {
		fmt.Printf("%6d  |  %s\n", k, v)
	}
}

//
// Each node periodically calls stabilize
// to learn about newly joined nodes
//
func (n *Node) stabilize() error {


}

//
// Each node periodically calls fix fingers to 
// make sure its finger table entries are correct
//
func (n *Node) fixFinger() error {
	// new nodes initialize their finger tables
	// existing nodes incorporate new nodes into their finger tables

}

//
// Each node periodically calls check predecessor
// to clear the node’s predecessor pointer if n.predecessor has failed
// to accept a new predecessor in notify
//
func (n *Node) checkPredecessor() error {


}