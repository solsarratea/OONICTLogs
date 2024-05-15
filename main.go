package main

import (
	"fmt"

	"./finder"
	"./submitter"
)

func main() {
	fmt.Println("Starting OONICTLogs...")

	go finder.Start()
	go submitter.Start()

	select {}
}
