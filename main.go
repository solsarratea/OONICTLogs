package main

import (
	"flag"
	"fmt"

	"github.com/solsarratea/OONICTLogs/finder"
	"github.com/solsarratea/OONICTLogs/submitter"
)

// Receives OS arguments and starts: finder and submitter processes
func main() {

	logChannel := make(chan string, 2)
	logChannel <- "Starting OONICTLogs..."

	find := flag.Bool("find", false, "Enable the finder process")
	submit := flag.Bool("submit", false, "Enable the submit process")
	flag.Parse()

	if *find {
		go finder.Start(logChannel)
	}

	if *submit {
		go submitter.Start(logChannel)
	}

	for msg := range logChannel {
		fmt.Println(msg)

	}

	select {}
}
