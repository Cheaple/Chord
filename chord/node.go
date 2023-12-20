package chord

import (
	"fmt"
	// "log"
	"math/big"
	"os"
	"time"

	"chord/utils"
	"chord/chord/rpc"
	"google.golang.org/grpc"
)



/* ******************************************************************************* *
 * ******************************** Basic Operations ***************************** */

//
// Create a new Chord n,
// and create or join a Chord ring
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
	node.FingerTable = node.makeNodeTable(M + 1)  // one more element for the node itself; real fingers start from index 1
	node.Predecessor = ""
	node.Successors = node.makeNodeTable(args.CntSuccessors + 1)  // one more element for the node itself; real successors start from index 1
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
		// node.create()
	}

	// Periodically stabilize
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

	// Periodically fix finger tables
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

	// Periodically checkes predecessor	
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

// //
// // Create a new Chord ring
// //
// func (n *Node) create() {
// 	n.Predecessor = ""
// 	// n.Predecessor = n.Address
// }

//
// Join an old Chord ring that contains a node at joinedAddress
//
func (n *Node) join(joinedAddress NodeAddress) error {
	n.Predecessor = ""
	
	succ, err := n.findSuccessorRPC(foo, n.Id)
	if err != nil {
		return err
	}
	n.successor = succ
	
	return nil
}

//
// Find an identifier's successor in the Chord ring
//
func (n *Node) findSuccessor(id *bigInt) (NodeEntry, error) {
	succ := n.Successors[0]
	succId := new(bigInt).SetBytes(succ.Identifier)
	if between(n.Id, id, succId, true) {
		return succ, nil
	}

	pred := n.closestPreceding(id)

	return rpc.findSuccessor(pred, id)
}

//
// Search the local table for the highest predecessor of id
//
func (n *Node) closestPreceding(id *bigInt) NodeEntry {
	for i := M; i > 0; i-- {
		prec := n.FingerTable[i]
		precId := new(bigInt).SetBytes(prec.Identifier)
		if between(n.Id, precId, id, false) {
			return prec
		}
	}
	return n.FingerTable[0]
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