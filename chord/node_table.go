package chord

// For Chord node's Finger Table & Successor List

import (
	// "math/big"
)

//
// Init NodeEntry
//
// func makeNodeEntry() NodeEntry {

// }


//
// Init Finger Table or Successor List
//
func (n *Node) makeNodeTable(size int) NodeTable {
	tbl := make([]NodeEntry, size)

	for i := range tbl {
		tbl[i].Id = n.Identifier.Bytes()
		tbl[i].Address = n.Address
	}

	return tbl
}