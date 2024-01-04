package chord

//
// File Operations for Chord
//

import (
	// "io/ioutil"
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
		if err := os.MkdirAll(nodeDir, os.ModePerm); err != nil {
			log.Fatalf("Error creating node folder: " + err.Error())
		} 	
	}

	// Upload directory (if already existing, delete & re-create)
	uploadDir := path.Join(nodeDir, "upload")
	if _, err := os.Stat(uploadDir); err == nil {
		if err := os.RemoveAll(uploadDir); err != nil {
			log.Fatalf("Error clearing upload folder: " + err.Error())
		}
	}
	if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating node upload folder: " + err.Error())
	}

	// Backup directory (if already existing, delete & re-create)
	backupDir := path.Join(nodeDir, "backup")
	if _, err := os.Stat(backupDir); err == nil {
		if err := os.RemoveAll(backupDir); err != nil {
			log.Fatalf("Error clearing upload folder: " + err.Error())
		}
	}
	if err := os.Mkdir(backupDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating node upload folder: " + err.Error())
	}

	// Certificate directory
	cerDir := path.Join(nodeDir, "certificate")
	if _, err := os.Stat(cerDir); err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(cerDir, os.ModePerm); err != nil {
			log.Fatalf("Error creating node upload folder: " + err.Error())
		}
	}

	n.baseDir = nodeDir
}

//
// Return the complete path of a given file
//
func (n *Node) getFilePath(fileName string) string {
	uploadDir := path.Join(n.baseDir, "upload")
	return path.Join(uploadDir, fileName)
}

//
// Return the complete path of a given backup file
//
func (n *Node) getBackupPath(fileName string) string {
	backupDir := path.Join(n.baseDir, "backup")
	return path.Join(backupDir, fileName)
}
