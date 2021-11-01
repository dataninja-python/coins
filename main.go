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
	// change := []int{1, 1, 1, 1, 1}
	// change := []int{1, 5, 1, 1, 1, 10, 15, 20, 100}
	// change := []int{6, 4, 5, 1, 1, 8, 9}
	// change := []int{}
	// change := []int{87}
	// change := []int{5, 6, 1, 1, 2, 3, 4, 9}
	// change := []int{5, 6, 1, 1, 2, 3, 43}
	// change := []int{1, 1}
	// change := []int{2}
	// change := []int{1}
	// change := []int{109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4}
	// change := []int{}
	// minChange := nonconstructible.NonConstructibleChange(change)
	// minChange2 := nonconstructible.NonConstructibleChange2(change)
	// fmt.Println(minChange2)
	// nonconstructible.WalkArray(change)
	// minChange3 := nonconstructible.TestArray2(change)
	minChange4 := nonconstructible.TestArray3(change)
	return minChange4
}

func main() {
	defer timeTrack(time.Now(), "coins1")
	fmt.Println("Minimum change not constructible: ", execute())
}
