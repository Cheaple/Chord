package chord

import (
	"fmt"
	// "log"
	"math/big"
	"os"
	"time"

	"chord/utils"
)



/* ******************************************************************************* *
 * ******************************** Basic Operations ***************************** */

//
// Create a new Chord n,
// and create or join a Chor ring
//
func MakeNode(args utils.Arguments) *Node {
	node := &Node{
		doneCh: make(chan struct{}),
	}

	node.Address = NodeAddress(fmt.Sprintf("%s:%d", args.Address, args.Port))
	if args.Identifier == "" {
		node.Identifier = hashString(string(node.Address))
	} else {
		node.Identifier = new(big.Int)
		node.Identifier.SetString(args.Identifier, 16)  // string of 16-bit int to big.Int
	}


	node.FingerTable = node.makeNodeTable(M)
	node.Successors = node.makeNodeTable(args.CntSuccessors)
	node.Bucket = make(map[Key]string)	



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
}

//
// Join an old Chord ring
//
func (n *Node) join(joined NodeAddress) error {
	n.Predecessor = ""




	return nil
}


/* ******************************************************************************* *
 * ******************************** Client Operations **************************** */

//
// Print current state
//
func (n *Node) Print() {
	fmt.Println("Node Address:", n.Address)
	fmt.Println("Node Identifier:", new(big.Int).SetBytes(n.Identifier.Bytes()))
	fmt.Println("Node Predecessor:", n.Predecessor)

	fmt.Println("------ Successor List ------")
	fmt.Println("Successors  |  Identifier  |  Address ")
	for i := 0; i < len(n.Successors); i++ {
		entry := n.Successors[i]
		id := new(big.Int).SetBytes(entry.Id)
		address := entry.Address
		fmt.Printf("%6d  |  %10d  |  %s\n", i, id, address)
	}
	
	fmt.Println("------ Finger Table ------")
	fmt.Println("Finger  |  Identifier  |  Address ")
	for i := 0; i < len(n.FingerTable); i++ {
		entry := n.FingerTable[i]
		fmt.Printf("%6d  |  %10d  |  %s\n", i, entry.Id, entry.Address)
	}

	fmt.Println("----- Buckets -----")
	fmt.Println("Key  |  Identifier  |  Address ")
	for k, v := range n.Bucket {
		fmt.Printf("%6d  |  %s\n", k, v)
	}
}


//
// Look up a key in the Chord ring
//
func (n *Node) LookUp(key string) (string, error) {


	return "", nil
}

//
// Download a file from the Chord ring
//
func (n *Node) Get(fileName string) (string, error) {

	return "", nil
}

//
// Store a file in the Chord ring
//
func (n *Node) Store(filePath string) (string, error) {


	return "", nil
}




/* ******************************************************************************* *
 * ****************************** Periodical Operations ************************** */

//
// Each node periodically calls stabilize
// to learn about newly joined nodes
//
func (n *Node) stabilize() error {


	return nil
}

//
// Each node periodically calls fix fingers to 
// make sure its finger table entries are correct
//
// paras:
// 	next: stores the index of the next finger to fix
//
func (n *Node) fixFinger(next int) int {
	// new nodes initialize their finger tables
	// existing nodes incorporate new nodes into their finger tables
	

	return 0
}

//
// Each node periodically calls check predecessor
// to clear the node’s predecessor pointer if n.predecessor has failed
// to accept a new predecessor in notify
//
func (n *Node) checkPredecessor() error {

	return nil
}