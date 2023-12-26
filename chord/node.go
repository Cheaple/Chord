package chord

import (
	"fmt"
	"log"
	"math/big"
	// "os"
	"time"

	"chord/utils"
)


//
//  Helper function for printng debugging logs
//
func (n *Node) DPrintf(str string, args ...interface{}) {
	if n.verbose != true {
		return
	}
	message := fmt.Sprintf(str, args...)
	fmt.Println()
	log.Printf(fmt.Sprintf("--- node-%d: %s ---", n.Id, message))
}


/* ******************************************************************************* *
 * ******************************** Basic Operations ***************************** */

//
// Create a new Chord n,
// and create or join a Chord ring
//
func NewNode(args utils.Arguments) *Node {
	node := &Node{
		doneCh: make(chan struct{}),
		verbose: args.Verbose,
	}

	node.Address = NodeAddress(fmt.Sprintf("%s:%d", args.Address, args.Port))
	if args.IdentifierStr == "" {
		node.Id = hashString(string(node.Address))
	} else {
		node.Id = new(big.Int)
		node.Id.SetString(args.IdentifierStr, 16)  // string of 16-bit int to big.Int
	}
	node.Entry = newNodeEntry(node.Id, node.Address)

	node.FingerTable = node.newNodeList(M + 1)  // one more element for the node itself; real fingers start from index 1
	node.Predecessor = &NodeEntry{}  // set empty node entry
	node.Successors = node.newNodeList(args.CntSuccessors + 1)  // one more element for the node itself; real successors start from index 1
	node.Bucket = make(map[Key]string)	

	// Join or Create a Chord ring
	if args.JoinAddress != "" {
		joinAddress := NodeAddress(fmt.Sprintf("%s:%d", args.JoinAddress, args.JoinPort))
		fmt.Println("Joining a Chord ring at node", joinAddress)
		err := node.joinChord(joinAddress)
		if err != nil {
			log.Fatal("Error joining the Chord ring:", err)
		}
	} else {
		fmt.Println("Creating a new Chord ring at node", node.Address)
		// node.create()
	}

	// Start RPC service to communicate with other Chord nodes
	node.StartRPCService()

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


//
// Join an old Chord ring that contains a node at joinedAddress
// Note: 
//   This function itself do not comminicate with other nodes at all.		
//   This node becomes a node of the Chord ring only after periodical stabilize().
//
func (n *Node) joinChord(joinedAddress NodeAddress) error {
	// Target Chord ring
	target := newNodeEntry(hashString(string(joinedAddress)), joinedAddress)

	n.Predecessor = &NodeEntry{}  // set empty node entry

	// Find the successor of this node in the target Chord ring
	succ, err := n.LocateRPC(target, n.Id)
	if err != nil {
		return err
	}
	n.Successors.set(1, succ)

	return nil
}

//
// Locate an identifier's successor in the Chord ring
// find_Successor() function in the paper
//
func (n *Node) locateSuccessor(id *big.Int) (*NodeEntry, error) {
	n.DPrintf("locateSuccessor(): id = %d", id)
	succ := n.Successors.get(1)
	succId := new(big.Int).SetBytes(succ.Identifier)
	if between(n.Id, id, succId, true) {
		// target id in (n, n.successor]
		return succ, nil
	}

	pred := n.closestPreceding(id)

	return n.LocateRPC(pred, id)
}

//
// Search the local table for the highest predecessor of id
// closest_preceding_node() function in the paper
//
func (n *Node) closestPreceding(id *big.Int) *NodeEntry {
	n.DPrintf("closestPreceding(): id = %d", id)
	for i := M; i > 0; i-- {
		prec := n.FingerTable.get(i)
		precId := new(big.Int).SetBytes(prec.Identifier)
		if between(n.Id, precId, id, false) {
			return prec
		}
	}
	return n.FingerTable.get(0)
}


/* ******************************************************************************* *
 * ****************************** Periodical Operations ************************** */

//
// Each node periodically calls stabilize
// to learn about newly joined nodes
//
func (n *Node) stabilize() {
	n.DPrintf("stabilize()")

	// Find the successor's current predecessor
	succ := n.Successors.get(1)
	succPred, err := n.GetPredecessorRPC(succ)
	if err != nil {
		fmt.Println("Error stabilizing:", err)
		return
	}
	
	// Update the successor
	if nodeBetweenOpen(n.Entry, succPred, succ) {
		// if succPred in (n, succ)
		n.Successors.set(1, succPred)
		// TODO: update successor list
	}

	// Notify the successor to update its predecessor
	_, err = n.NotifyRPC(n.Successors.get(1))
	if err != nil {
		fmt.Println("Error notifying:", err)
	}
}

//
// Each node periodically calls fix fingers to 
// make sure its finger table entries are correct
// new nodes initialize their finger tables
// existing nodes incorporate new nodes into their finger tables
//
// paras:
// 	next: stores the index of the next finger to fix
//
func (n *Node) fixFinger(next int) int {
	next = next % M + 1  // next in [1, M]

	// next finger's id = n.id + 2 ^ next
	cur := new(big.Int).Set(n.Id)
	add := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(next) - 1), nil)
	nextId := new(big.Int).Add(cur, add)
	nextId = new(big.Int).Mod(nextId, hashMod)
	n.DPrintf("fixFinger(): next id = %d", nextId)

	// Find a successor node that stores the next finger id
	finger, err := n.locateSuccessor(nextId)
	if err != nil || finger == nil {
		fmt.Println("Error fixing finger table:", err)
	}

	// Update finger entry
	n.FingerTable.set(next, finger)
	n.DPrintf("fixFinger(): finger[%d] = %+v", next, finger)
	return next
}

//
// Each node periodically calls check predecessor
// to clear the node’s predecessor pointer if n.predecessor has failed
// to accept a new predecessor in notify
//
func (n *Node) checkPredecessor() error {
	if n.Predecessor.empty() {
		return nil
	}
	_, err := n.CheckRPC(n.Predecessor)
	if err != nil {
		n.DPrintf("checkPredecessor(): set n.predecessor = nil")
		n.Predecessor = &NodeEntry{}  // set empty node entry
	}
	return nil
}