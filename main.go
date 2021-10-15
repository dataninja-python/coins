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
	defer timeTrack(time.Now(), "coins1")
	change := []int{5, 7, 1, 1, 2, 3, 22}
	minChange := nonconstructible.NonConstructibleChange(change)
	return minChange
}

func main() {
	fmt.Println(execute())
}
