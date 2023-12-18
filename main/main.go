package main

import (
	"bufio"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
	"time"

	"chord/chord"
	"chord/utils"
)



func ConnHandler(listener net.Listener, node *chord.Node) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept failed:", err.Error())
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func main() {
	// Parse command line arguments
	args, err := utils.ParseCmdArgs()
	if err != nil {
		fmt.Println("Error parsing arguments: ", err)
		os.Exit(1)
	}

	node := chord.NewNode(args)
	rpc.Register(node)

	////************************** Config network info and deployment *****************************////
	IPAddr := fmt.Sprintf("%s:%d", args.Address, args.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", IPAddr)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	// Listen to the address
	listener, err := net.Listen("tcp", tcpAddr.String())
	if err != nil {
		fmt.Println("ListenTCP failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Local node listening on ", tcpAddr)

	// Use a separate goroutine to accept connection
	go ConnHandler(listener, node)

	////************************* Chord node initialization operation *****************************////
	if args.JoinAddress != "Unspecified" {  // Join exsiting chord
		RemoteAddr := fmt.Sprintf("%s:%d", args.JoinAddress, args.JoinPort)
		fmt.Println("Connecting to the remote node..." + RemoteAddr)
		err := node.JoinChord(chord.NodeAddress(RemoteAddr))
		if err != nil {
			fmt.Println("Join RPC call failed")
			os.Exit(1)
		} else {
			fmt.Println("Join RPC call success")
		}
	} else {  // Create new chord
		node.CreateChord()
	}

	
	////************************************* Chordshell *****************************************////
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s@Chord: node-%s > ", node.Addr, node.Identifier)
		input, _ := reader.ReadString('\n')
		cmdArgs := strings.Fields(input)
		if len(cmdArgs) < 1 {
			continue
		}
		cmd := strings.ToUpper(strings.TrimSpace(cmdArgs[0]))

		if cmd == "PRINTSTATE" || cmd == "PS" {
			fmt.Println("-------------- Current Node State --------------")
			node.PrintState()
			continue

		} else if cmd == "LOOKUP" || cmd == "L" {
			if len(cmdArgs) < 2 {
				fmt.Printf("Invalid command")
				fmt.Println("usage: LOOKUP <filename>")
				continue
			}

			// Look up for a filename in the Chord ring
			filename := cmdArgs[1]
			resultAddr, err := chord.ClientLookUp(filename, node)
			if err != nil {
				fmt.Println("Error looking up file:", err)
				continue
			}
			fmt.Printf("The file should be in Node (%s)\n", resultAddr)

			// Check if the file is stored in the node
			reply := chord.CheckFileExistRPCReply{}
			err = chord.ChordCall(resultAddr, "Node.CheckFileExistRPC", filename, &reply)
			if err != nil {
				fmt.Println("Error checking file's existence:", err)
				continue
			}

			if reply.Exist {
				// Get the address of the node that stores the file
				var getNameRPCReply chord.GetNameRPCReply
				err = chord.ChordCall(resultAddr, "Node.GetNameRPC", "", &getNameRPCReply)
				if err != nil {
					fmt.Println("Node.GetNameRPC() RPC call failed in main()")
				} else {
					fmt.Println("The file is stored at", getNameRPCReply.Name)
				}
			} else {
				fmt.Println("The file is not stored in the node")
			}

		} else if cmd == "GET" || cmd == "G" {
			if len(cmdArgs) < 2 {
				fmt.Printf("Invalid command. ")
				fmt.Println("Usage: GET <filename>")
				continue
			}
			
			// Download a file from the Chord ring
			fileName := cmdArgs[1]
			err := chord.ClientGetFile(fileName, node)
			if err != nil {
				continue
			}
			fmt.Println("GET Success")

		} else if cmd == "STOREFILE" || cmd == "S" {
			if len(cmdArgs) < 2 {
				fmt.Printf("Invalid command. ")
				fmt.Println("Usage: STOREFILE <filepath>")
				continue
			}
			
			// Upload a local file to the Chord ring
			fileName := cmdArgs[1]
			err := chord.ClientStoreFile(fileName, node)
			if err != nil {
				fmt.Println("Error storing file:", err)
				continue
			} 
			fmt.Println("STOREFILE Success")

		} else if cmd == "QUIT" || cmd == "Q" {
			// Quit the program
			os.Exit(0)

		} else {
			fmt.Println("Invalid command")
		}
	}
}