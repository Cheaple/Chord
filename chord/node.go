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
func (n *Node) DPrintf(args ...interface{}) {
	if n.verbose != true {
		return
	}
	message := fmt.Sprint(args...)
	log.Println(fmt.Sprintf("----- node-%d () %s -----", n.Id, n.Address, message))
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

	node.FingerTable = node.newNodeTable(M + 1)  // one more element for the node itself; real fingers start from index 1
	node.Predecessor = &NodeEntry{}  // empty
	node.Successors = node.newNodeTable(args.CntSuccessors + 1)  // one more element for the node itself; real successors start from index 1
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
//   This function is different from the paper Fig.6. join(). 
//   This function itself do not comminicate with other nodes at all.		
//   This node becomes a node of the Chord ring only after periodical stabilize().
func (n *Node) joinChord(joinedAddress NodeAddress) error {
	n.Predecessor = &NodeEntry{}  // empty
	n.Successors[1] = newNodeEntry(hashString(string(joinedAddress)), joinedAddress)
	return nil
}

//
// Find an identifier's successor in the Chord ring
//
func (n *Node) findSuccessor(id *big.Int) (*NodeEntry, error) {
	n.DPrintf("findSuccessor(): id = %d", id)
	succ := n.Successors[0]
	succId := new(big.Int).SetBytes(succ.Identifier)
	if between(n.Id, id, succId, true) {
		return succ, nil
	}

	pred := n.closestPreceding(id)

	return n.findSuccessorRPC(pred, id)
}

//
// Search the local table for the highest predecessor of id
//
func (n *Node) closestPreceding(id *big.Int) *NodeEntry {
	n.DPrintf("closestPreceding(): id = %d", id)
	for i := M; i > 0; i-- {
		prec := n.FingerTable[i]
		precId := new(big.Int).SetBytes(prec.Identifier)
		if between(n.Id, precId, id, false) {
			return prec
		}
	}
	return n.FingerTable[0]
}


/* ******************************************************************************* *
 * ****************************** Periodical Operations ************************** */

//
// Each node periodically calls stabilize
// to learn about newly joined nodes
//
func (n *Node) stabilize() {
	succ := n.Successors[1]
	succPred, err := n.findPredecessorRPC(succ)
	if err != nil {
		fmt.Println("Error stabilizing:", err)
		return
	}
	if nodeBetweenOpen(n.Entry, succPred, succ) {
		// if succPred in (n, succ)
		n.Successors[1] = succPred
	}
	// TODO: notify
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