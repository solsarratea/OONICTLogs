package submitter

import (
	"fmt"
	"time"
)

func Start() {
	fmt.Println("Starting Submitter")

	for {
		fmt.Println("Submitter running...")
		time.Sleep(5 * time.Second)
	}
}
