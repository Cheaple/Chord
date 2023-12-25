package chord

/*
 * For user operations
 * e.g. LookUp, StoreFile, GetFile, Quit
 */

import (
	"fmt"
	// "log"
	"math/big"
	// "os"
)


/* ******************************************************************************* *
 * ********************************* User Operations ***************************** */

//
// Print current state
//
func (n *Node) Print() {
	fmt.Println("Node Address:", n.Address)
	fmt.Println("Node Identifier:", new(big.Int).SetBytes(n.Id.Bytes()))
	if n.Predecessor.empty() {
		fmt.Println("Node Predecessor: nil")
	} else {
		fmt.Println("Node Predecessor:", new(big.Int).SetBytes(n.Predecessor.Identifier))
	}

	fmt.Println("------ Successor List ------")
	fmt.Println("Successors  |  Identifier  |  Address ")
	for i := 1; i < len(n.Successors); i++ {
		entry := n.Successors[i]
		if entry.empty() {
			fmt.Printf("%10d  |              |\n", i)
			continue
		}
		id := new(big.Int).SetBytes(entry.Identifier)
		address := entry.Address
		fmt.Printf("%10d  |  %10d  |  %s\n", i, id, address)
	}
	
	fmt.Println("------ Finger Table ------")
	fmt.Println("    Finger  |  Identifier  |  Address ")
	for i := 1; i < len(n.FingerTable); i++ {
		entry := n.FingerTable[i]
		if entry.empty() {
			fmt.Printf("%10d  |              |\n", i)
			continue
		}
		id := new(big.Int).SetBytes(entry.Identifier)
		address := entry.Address
		fmt.Printf("%10d  |  %10d  |  %s\n", i, id, address)
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