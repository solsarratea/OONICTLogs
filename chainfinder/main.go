package chainfinder

import (
	"fmt"
	"time"
)

func Start() {
	fmt.Println("Starting Chainfinder")

	for {
		fmt.Println("Chainfinder running...")
		time.Sleep(5 * time.Second)
	}
}
