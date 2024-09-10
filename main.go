package main

import (
	"fmt"
	"os"

	"github.com/solsarratea/OONICTLogs/finder"
	"github.com/solsarratea/OONICTLogs/submitter"
)

// Receives OS arguments and starts: finder and submitter processes
func main() {
	fmt.Println("Starting OONICTLogs...")

	ENABLE_FINDER := false
	ENABLE_SUBMITTER := false

	if len(os.Args) == 1 || os.Args[1] == "all" {
		ENABLE_FINDER = true
		ENABLE_SUBMITTER = true
	} else {
		for _, arg := range os.Args[1:] {
			switch arg {
			case "find":
				ENABLE_FINDER = true
			case "submit":
				ENABLE_SUBMITTER = true
			}
		}
	}

	if ENABLE_FINDER {
		go finder.Start()
	}

	if ENABLE_SUBMITTER {
		go submitter.Start()
	}

	select {}
}
