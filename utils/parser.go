package utils

// Parse command line arguments

import (
	"errors"
	"flag"
	"net"
	"regexp"
)

type Arguments struct {
	Address     	string	// IP address that the Chord client will bind to
	Port        	int     // Port that the Chord client will bind to and listen on
	JoinAddress 	string	// IP address of the machine running a Chord node. The Chord client will join this node’s ring.
	JoinPort    	int     // Port that an existing Chord node is bound to and listening on. The Chord client will join this node’s ring.
	Stabilize_t 	int     // Time in milliseconds between invocations of stabilize
	FixFinger_t 	int     // Time in milliseconds between invocations of fix_finger.
	CheckPred_t 	int     // Time in milliseconds between invocations of check_predecessor
	CntSuccessors  	int		// Number of successors maintained by the Chord client
	Identifier  	string  // Identifier
}

//
// Parse command line arguments
//
func ParseCmdArgs() (Arguments, error) {
	var a		 	string 
	var p 			int    	
	var ja 			string	
	var jp 			int    	
	var ts 			int
	var tff 		int
	var tcp 		int   		
	var r 			int     	
	var i 			string  	

	// Parse command line arguments
	flag.StringVar(&a, "a", "localhost", "IP address that the Chord client will bind to, as well as advertise to other nodes. Represented as an ASCII string (e.g., 128.8.126.63). Must be specified.")
	flag.IntVar(&p, "p", 8000, "Port that the Chord client will bind to and listen on. Represented as a base-10 integer. Must be specified.")
	flag.StringVar(&ja, "ja", "", "IP address of the machine running a Chord node. The Chord client will join this node’s ring. Represented as an ASCII string (e.g., 128.8.126.63). Must be specified if --jp is specified.")
	flag.IntVar(&jp, "jp", 8000, "Port that an existing Chord node is bound to and listening on. The Chord client will join this node’s ring. Represented as a base-10 integer. Must be specified if --ja is specified.")
	flag.IntVar(&ts, "ts", 3000, "Time in milliseconds between invocations of ‘stabilize’. Represented as a base-10 integer. Must be specified, with a value in the range of [1,60000].")
	flag.IntVar(&tff, "tff", 1000, "Time in milliseconds between invocations of ‘fix fingers’. Represented as a base-10 integer. Must be specified, with a value in the range of [1,60000].")
	flag.IntVar(&tcp, "tcp", 3000, "Time in milliseconds between invocations of ‘check predecessor’. Represented as a base-10 integer. Must be specified, with a value in the range of [1,60000].")
	flag.IntVar(&r, "r", 3, "Number of successors maintained by the Chord client. Represented as a base-10 integer. Must be specified, with a value in the range of [1,32].")
	flag.StringVar(&i, "i", "", "Identifier (ID) assigned to the Chord client which will override the ID computed by the SHA1 sum of the client’s IP address and port number. Represented as a string of 40 characters matching [0-9a-fA-F]. Optional parameter.")
	flag.Parse()

	args := Arguments{
		Address:     	string(a),
		Port:        	p,
		JoinAddress: 	string(ja),
		JoinPort:    	jp,
		Stabilize_t: 	ts,
		FixFinger_t: 	tff,
		CheckPred_t: 	tcp,
		CntSuccessors:  r,
		Identifier:  	i,
	}
	err := validateArgs(args)
	
	return args, err
}

//
// Validate command line arguments
//
func validateArgs(args Arguments) error {
	if net.ParseIP(string(args.Address)) == nil && args.Address != "localhost" {
		return errors.New("Invalid argument -a (Ip address for the current node)")
	}

	if args.Port < 1024 || args.Port > 65535 {
		return errors.New("Invalid argument -p (port for the current node)")
	}

	if args.Stabilize_t < 1 || args.Stabilize_t > 60000 {
		return errors.New("Invalid argument -ts")
	}
	
	if args.FixFinger_t < 1 || args.FixFinger_t > 60000 {
		return errors.New("Invalid argument -tff")
	}
	
	if args.CheckPred_t < 1 || args.CheckPred_t > 60000 {
		return errors.New("Invalid argument -tcp")
	}

	if args.CntSuccessors < 1 || args.CntSuccessors > 32 {
		return errors.New("Invalid argument -r")
	}

	if args.Identifier != "" {
		matched, err := regexp.MatchString("[0-9a-fA-F]*", args.Identifier)
		if err != nil || !matched {
			return errors.New("Invalid argument -i (identifier)")
		}
	}

	if args.JoinAddress != "" {
		// -ja is specified, validate -ja & -jp
		if net.ParseIP(string(args.JoinAddress)) == nil && args.JoinAddress == "localhost" {
			return errors.New("Invalid argument -ja (port for the joined node)")
		}
		if args.JoinPort < 1024 || args.JoinPort > 65535 {
			return errors.New("Invalid argument -jp (port for the joined node)")
		}
	}

	return nil
}
