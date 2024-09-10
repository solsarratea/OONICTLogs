package main

import (
	"flag"
	"fmt"

	"github.com/solsarratea/OONICTLogs/finder"
	"github.com/solsarratea/OONICTLogs/submitter"
)

// Receives OS arguments and starts: finder and submitter processes
func main() {
	find := flag.Bool("find", false, "Enable the finder process")
	submit := flag.Bool("submit", false, "Enable the submit process")
	flag.Parse()

	fmt.Println("Starting OONICTLogs...")
	if *find {
		go finder.Start()
	}

	if *submit {
		go submitter.Start()
	}

	select {}
}
