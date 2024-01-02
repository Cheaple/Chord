package chord

//
// File Operations for Chord
//

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

const rootDir = "./data"

//
// Initialize data storage for this node, using local file system
//
func (n *Node) startDataStore() {
	// Root data directory
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		// Create a root data folder
		err := os.MkdirAll(rootDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating root data folder: " + err.Error())
		} 
	}

	// Node directory
	nodeDir := path.Join(rootDir, string(n.Address))
	if _, err := os.Stat(nodeDir); os.IsNotExist(err) {
		// Create a new node folder
		err := os.MkdirAll(nodeDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating node folder: " + err.Error())
		} 
		err = os.Mkdir(path.Join(nodeDir, "upload"), os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating node upload folder: " + err.Error())
		}
		err = os.Mkdir(path.Join(nodeDir, "certificate"), os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating node certificate folder: " + err.Error())
		}
	} else { 
		// Read a existing folder
		// Read node buckets from local files
		files, err := ioutil.ReadDir(path.Join(nodeDir, "upload"))
		if err != nil {
			log.Fatalf("Error reading node f: " + err.Error())
		}
		for _, f := range files {
			// Store file name in bucket
			name := f.Name()
			n.Bucket[name] = 1
		}
	}

	n.baseDir = nodeDir
}


func (n *Node) getFilePath(fileName string) string {
	uploadDir := path.Join(n.baseDir, "upload")
	return path.Join(uploadDir, fileName)
}
