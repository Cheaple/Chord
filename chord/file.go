package chord

//
// File Operations for Chord
//

import (
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const rootDir = "./data"

//
// Initialize data storage for this node, using local file system
//
func (n *Node) StartDataStore() {
	nodeDir := path.Join(rootDir, string(n.Address))

	if _, err := os.Stat(nodeDir); os.IsNotExist(err) {
		// Create a new node folder
		err := os.MkdirAll(nodeDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating node folder: " + err.Error())
		} 
		err = os.Mkdir(path.Join(nodeDir, "upload"), os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating node folder: " + err.Error())
		}
	} else { 
		// Read a existing folder
		// Read node buckets from local files
		files, err := ioutil.ReadDir(path.Join(nodeDir, "upload"))
		if err != nil {
			log.Fatalf("Error reading node files: " + err.Error())
		}
		for _, f := range files {
			// Store file name in bucket
			name := f.Name()
			n.Bucket[name] = 1
		}
	}
}
