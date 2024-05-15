package main

import (
	"fmt"

	"./chainfinder"
	"./submitter"
)

func main() {
	fmt.Println("Starting OONICTLogs...")

	go chainfinder.Start()
	go submitter.Start()

	select {}
}
