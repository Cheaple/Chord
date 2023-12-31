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
	// "path/filepath"
)

const rootDir = "./data"

//
// Initialize data storage for this node, using local file system
//
func (n *Node) startDataStore() {
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

//
// Get a file from the local data store  
//
// func (n *Node) getFile(filePath string) error {
// 	// Open file and pack into fileRPC
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		// fmt.Println("Error opening file:", err)
// 		return err
// 	}
// 	defer file.Close()

// }

func (n *Node) getFilePath(fileName string) string {
	nodeDir := path.Join(rootDir, string(n.Address))
	uploadDir := path.Join(nodeDir, "upload")
	return path.Join(uploadDir, fileName)
}
