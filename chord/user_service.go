package chord

/*
 * For user operations
 * e.g. LookUp, StoreFile, GetFile, Quit
 */

import (
	"fmt"
	// "log"
	"math/big"
	"path/filepath"
	"os"
)


/* ******************************************************************************* *
 * ********************************* User Operations ***************************** */

//
// Print current state
//
func (n *Node) Print() {
	fmt.Println("Node Address:", n.Address)
	fmt.Println("Node Address for TLS:", n.TlsAddress)
	fmt.Println("Node Identifier:", new(big.Int).SetBytes(n.Id.Bytes()))
	if n.Predecessor.empty() {
		fmt.Println("Node Predecessor: nil")
	} else {
		fmt.Println("Node Predecessor:", new(big.Int).SetBytes(n.Predecessor.Identifier))
	}

	fmt.Println("------ Successor List ------")
	fmt.Println("Successors  |  Identifier  |  Address ")
	for i := 1; i < len(n.Successors.Entries); i++ {
		entry := n.Successors.get(i)
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
	for i := 1; i < M + 1; i++ {
		entry := n.FingerTable.get(i)
		if entry.empty() {
			fmt.Printf("%10d  |              |\n", i)
			continue
		}
		id := new(big.Int).SetBytes(entry.Identifier)
		address := entry.Address
		fmt.Printf("%10d  |  %10d  |  %s\n", i, id, address)
	}

	fmt.Println("----- Buckets -----")
	fmt.Println("       Key  |  Indentifier")
	for k, _ := range n.Bucket {
		id := hashString(k)
		fmt.Printf("%10s  |  %10d\n", k, id)
	}
}


//
// Look up a key in the Chord ring
//
func (n *Node) LookUp(key string) (*NodeEntry, error) {
	targetNode, err := n.lookup(key)
	if err != nil {
		return nil, err
	}
	ifExist, err := n.CheckKeyRPC(targetNode, key)
	if err != nil || ifExist == false {
		return targetNode, fmt.Errorf("File not found")
	}
	return targetNode, nil
}

//
// Download a file from the Chord ring
//
func (n *Node) Get(fileName string) (bool, error) {
	// Locate a node which stores the file
	targetNode, err := n.lookup(fileName)
	if err != nil {
		return true, err
	}
	n.DPrintf("The file should be stored in %s\n", targetNode.ToString())

	// Download the file from the target node
	return n.DownloadFileRPC(targetNode, fileName)
}

//
// Store a file in the Chord ring
//
func (n *Node) Store(filePath string) (bool, error) {
	// Check file existence
	if _, err := os.Stat(filePath); err != nil {
		return false, fmt.Errorf("Error checking file '%s'.", filePath)
	}

	// Locate a node to store the file
	fileName := filepath.Base(filePath)
	targetNode, err := n.lookup(fileName)
	if err != nil {
		return true, err
	}
	n.DPrintf("The file will be stored in %s\n", targetNode.ToString())

	// Transfer the file to the target node
	return n.UploadFileRPC(targetNode, filePath)
}