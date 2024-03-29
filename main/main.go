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
			res, err := node.LookUp(filename)
			if err != nil {
				if res != nil {
					fmt.Printf("File not found in %s\n", res.ToString())
				} else {
					fmt.Printf("Error locating file: %s\n", err)
				}
				continue
			}
			fmt.Printf("File found in %s\n", res.ToString())

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
				fmt.Println("GET Failure:", err)
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
				fmt.Println("STORE Failure:", err)
				continue
			} 
			fmt.Println("STORE Success")

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
		log.Fatal("Error parsing arguments: ", err)
	}

	node := chord.NewNode(args)
	runChordClient(node)	
}