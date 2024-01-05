package chord

import (
	"fmt"
	"log"
	"math/big"
	"os"
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
		lenSuccessors: args.CntSuccessors + 1,  // one more element for the node itself
		doneCh: make(chan struct{}),
		verbose: args.Verbose,
	}

	node.IP = args.Address
	node.Address = NodeAddress(fmt.Sprintf("%s:%d", args.Address, args.Port))
	node.TlsAddress = NodeAddress(fmt.Sprintf("%s:%d", args.Address, args.Port + 1))
	if args.IdentifierStr == "" {
		node.Id = hashString(string(node.Address))
	} else {
		node.Id = new(big.Int)
		node.Id.SetString(args.IdentifierStr, 16)  // string of 16-bit int to big.Int
	}
	node.Entry = newNodeEntry(node.Id, node.Address, node.TlsAddress)

	node.FingerTable = node.newNodeList(M + 1)  // one more element for the node itself; real fingers start from index 1
	node.Predecessor = &NodeEntry{}  // set empty node entry
	node.Successors = node.newNodeList(node.lenSuccessors)  // one more element for the node itself; real successors start from index 1
	node.Bucket = make(map[string]*big.Int)
	node.Backup = make(map[string]*big.Int)

	// Start data store
	node.startDataStore()

	// Start TLS configuration
	node.startTlsConfig()

	// Start RPC service to communicate with other Chord nodes
	node.startRPCService()

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
		next := 1
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

	// Periodically backup local keys in its first successor
	go func() {
		ticker := time.NewTicker(time.Duration(args.CheckPredTime) * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				node.backup()
			case <-node.doneCh:
				ticker.Stop()
				return
			}
		}
	}()

	// Periodically send local keys to the predecessor
	go func() {
		ticker := time.NewTicker(time.Duration(args.CheckPredTime) * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				node.transferKeys()
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
	target := newNodeEntry(hashString(string(joinedAddress)), joinedAddress, "")

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
		// if target id in (n, n.successor]
		return succ, nil
	}

	// Find the closest available preceding entry in the finger table
	startIdx := M
	for startIdx >= 1 {
		pred, idx := n.closestPreceding(id, startIdx)
		if nodeEqual(pred, n.Entry) {
			return pred, nil
		}
		res, err := n.LocateRPC(pred, id)
		if err == nil {
			return res, nil
		}

		// If a node fails during the find successor procedure, the lookup
		// proceeds, after a timeout, by trying the next best predecessor
		// among the nodes in the finger table and the successor list.
		startIdx = idx - 1
	}
	
	return nil, fmt.Errorf("Cannot locate target id: %d", id)
}

//
// Locate a Key's successor in the Chord ring
//
func (n *Node) lookup(key string) (*NodeEntry, error) {
	id := hashString(key)
	return n.locateSuccessor(id)
}

//
// Search the local table for the highest predecessor of id,
// starting from the given index of the finger table
// closest_preceding_node() function in the paper
//
func (n *Node) closestPreceding(id *big.Int, startIdx int) (*NodeEntry, int) {
	n.DPrintf("closestPreceding(): id = %d", id)
	n.fingerMu.RLock()
	defer n.fingerMu.RUnlock()
	for i := startIdx; i > 0; i-- {
		prec := n.FingerTable.get(i)
		precId := new(big.Int).SetBytes(prec.Identifier)
		if between(n.Id, precId, id, false) {
			return prec, i
		}
	}
	return n.FingerTable.get(0), 0
}


/* ******************************************************************************* *
 * ****************************** Periodical Operations ************************** */

//
// Each node periodically calls stabilize
// to update successor list and learn about newly joined nodes
//
func (n *Node) stabilize() {
	n.DPrintf("stabilize()")

	// Update successor list
	n.succMu.Lock()
	for i := 1; i < n.lenSuccessors; i++ {
		succ := n.Successors.get(i)
		succSucc, err := n.GetSuccessorListRPC(succ)  // successor's successor list
		if err != nil {
			// this successor fails, go to the next successor
			continue
		}

		// reconciles its list with its successor succ 
		// by copying succ’s successor list, 
		// removing its last entry, and prepending succ to it.
		n.Successors.set(1, succ)
		for i := 1; i < n.lenSuccessors - 1; i++ {
			n.Successors.set(i + 1, succSucc.get(i))
		}
		n.DPrintf("stabilize(): update Successor List: %+v", n.Successors)
		break
	}
	n.succMu.Unlock()

	// Find the successor's current predecessor
	succ := n.Successors.get(1)
	succPred, err := n.GetPredecessorRPC(succ)
	if err != nil {
		fmt.Println("Error stabilizing:", err)
		return
	}
	
	// Update successor list if newly joining node
	n.succMu.Lock()
	if nodeBetweenOpen(n.Entry, succPred, succ) {
		// if succPred in (n, succ)
		for i := n.lenSuccessors - 1; i > 1; i-- {
			n.Successors.set(i, n.Successors.get(i - 1))
		}
		n.Successors.set(1, succPred)
		n.DPrintf("stabilize(): update Successor List: %+v", n.Successors)
	}
	n.succMu.Unlock()

	// Notify the successor to update its predecessor
	n.succMu.RLock()
	defer n.succMu.RUnlock()
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
	// next finger's id = n.id + 2 ^ next
	cur := new(big.Int).Set(n.Id)
	add := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(next) - 1), nil)
	nextId := new(big.Int).Add(cur, add)
	nextId = new(big.Int).Mod(nextId, hashMod)
	// n.DPrintf("fixFinger(): next id = %d", nextId)

	// Find a successor node that stores the next finger id
	finger, err := n.locateSuccessor(nextId)
	if err != nil || finger == nil {
		fmt.Println("Error fixing finger table:", err)
		return next % M + 1
	}

	// Update finger entry
	n.fingerMu.Lock()
	defer n.fingerMu.Unlock()
	n.FingerTable.set(next, finger)
	// n.DPrintf("fixFinger(): finger[%d] = %+v", next, finger)
	return next % M + 1  // next in [1, M]
}

//
// Each node periodically calls check predecessor
// to clear the node’s predecessor pointer if n.predecessor has failed
// to accept a new predecessor in notify
//
func (n *Node) checkPredecessor() error {
	n.predMu.RLock()
	if n.Predecessor.empty() {
		n.predMu.RUnlock()
		return nil
	}
	_, err := n.CheckRPC(n.Predecessor)
	n.predMu.RUnlock()

	// Deal with predecessor failure
	if err != nil {
		n.DPrintf("checkPredecessor(): set n.predecessor = nil")
		n.predMu.Lock()
		n.Predecessor = &NodeEntry{}  // set empty node entry
		n.predMu.Unlock()

		// move backup files of the failed predecessor to the local buckets
		n.backupMu.RLock()
		keysToRecover := make([]string, 0)
		for key, _ := range n.Backup {
			keysToRecover = append(keysToRecover, key)
		}
		n.backupMu.RUnlock()

		for _, key := range keysToRecover {
			err := os.Rename(n.getBackupPath(key), n.getFilePath(key))
			if err != nil {
				continue
			}
			delete(n.Backup, key)
			n.Bucket[key] = hashString(key)
			n.DPrintf("checkPredecessor(): recover backup file %s", key)
		}
	}

	return nil
}

//
// Each node periodically calls backup()
// to send local files to its first successor
//
func (n *Node) backup() error {
	n.succMu.RLock()
	defer n.succMu.RUnlock()
	if nodeEqual(n.Successors.get(1), n.Entry) {
		// if the first successor is the node itself, do not backup
		return nil
	}

	// call UploadRPC to send local files to the first successor to backup
	n.bucketMu.RLock()
	defer n.bucketMu.RUnlock()
	for key, _ := range n.Bucket {
		n.UploadFileRPC(n.Successors.get(1), n.getFilePath(key), true)
	}
	
	return nil
}

//
// Each node periodically calls transferKeys()
// to transfer parts of local keys to the predecessor
//
func (n *Node) transferKeys() error {
	n.predMu.RLock()
	defer n.predMu.RUnlock()
	if n.Predecessor.empty() {
		return nil
	}
	pred := n.Predecessor
	predId := new(big.Int).SetBytes(pred.Identifier)

	// Check & transfer keys that should not belong to the current node
	n.bucketMu.Lock()
	defer n.bucketMu.Unlock()
	keysToDelete := make([]string, 0)
	for key, id := range n.Bucket {
		if between(predId, id, n.Id, true) {
			// if key in (n.predecessor, n], no need to transfer
			continue  // skip this key
		}

		// transfer this key to the predecessor
		keysToDelete = append(keysToDelete, key)
		_, err := n.UploadFileRPC(n.Predecessor, n.getFilePath(key), false)
		if err != nil {
			continue
		}
	}

	// Remove transfered keys
	for _, key := range keysToDelete {
		err := os.Remove(n.getFilePath(key))
		if err != nil {
			continue
		}
		delete(n.Bucket, key)
		n.DPrintf("transferKeys(): tranfer local file '%s' to the predecessor %s", key, n.Predecessor.ToString())
	}
	
	return nil
}