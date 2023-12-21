package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"chord/chord"
	"chord/utils"
)

//
// Chord Client Shell
//
func runChordClient(node *chord.Node) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s@Chord: node-%s > ", node.Address, node.Id)
		input, _ := reader.ReadString('\n')
		cmdArgs := strings.Fields(input)
		if len(cmdArgs) < 1 {
			continue
		}
		cmd := strings.ToUpper(strings.TrimSpace(cmdArgs[0]))

		if cmd == "PRINTSTATE" || cmd == "PS" {
			node.Print()
			continue

		} else if cmd == "LOOKUP" || cmd == "L" {
			if len(cmdArgs) < 2 {
				fmt.Printf("Invalid command")
				fmt.Println("usage: LOOKUP <filename>")
				continue
			}

			// Look up for a filename in the Chord ring
			filename := cmdArgs[1]
			resultAddr, err := node.LookUp(filename)
			if err != nil {
				fmt.Println("Error looking up file:", err)
				continue
			}
			fmt.Printf("The file should be in Node (%s)\n", resultAddr)

			// TODO: check file existence

		} else if cmd == "GETFILE" || cmd == "GET" || cmd == "G" {
			if len(cmdArgs) < 2 {
				fmt.Printf("Invalid command. ")
				fmt.Println("Usage: GET <filename>")
				continue
			}
			
			// Download a file from the Chord ring
			fileName := cmdArgs[1]
			_, err := node.Get(fileName)
			if err != nil {
				continue
			}
			fmt.Println("GET Success")

		} else if cmd == "STOREFILE" || cmd == "STORE" || cmd == "S" {
			if len(cmdArgs) < 2 {
				fmt.Printf("Invalid command. ")
				fmt.Println("Usage: STOREFILE <filepath>")
				continue
			}
			
			// Upload a local file to the Chord ring
			fileName := cmdArgs[1]
			_, err := node.Store(fileName)
			if err != nil {
				fmt.Println("Error storing file:", err)
				continue
			} 
			fmt.Println("STOREFILE Success")

		} else if cmd == "QUIT" || cmd == "Q" {
			log.Println("Exit Chord Ring!")
			os.Exit(0)

		} else {
			fmt.Println("Invalid command!")
		}
	}

}

func main() {
	// Parse command line arguments
	args, err := utils.ParseCmdArgs()
	if err != nil {
		fmt.Println("Error parsing arguments: ", err)
		os.Exit(1)
	}

	node := chord.MakeNode(args)
	runChordClient(node)	
}