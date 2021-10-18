package main

import (
	"change/nonconstructible"
	"fmt"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func execute() int {
	change := []int{5, 7, 1, 1, 2, 3, 22}
	// change := []int{}
	minChange := nonconstructible.NonConstructibleChange(change)
	return minChange
}

func main() {
	defer timeTrack(time.Now(), "coins1")
	fmt.Println("Minimum Number: ", execute())
}
