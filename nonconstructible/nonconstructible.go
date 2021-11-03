package nonconstructible

import (
	"fmt"
	"sort"
)

func NonConstructibleChange(coins []int) int {

	fmt.Println("Initial coin slice: ", coins)
	// sort array
	sort.Ints(coins)
	fmt.Println("Sorted coin slice: ", coins)

	// set up a variable to track the change you can create
	change := 0
	fmt.Println("Starting change: ", change)
	for _, c := range coins {
		if c > change+1 {
			fmt.Println("returning early: ", change+1)
			return change + 1
		}
		change += c
	}
	fmt.Println("Final change: ", change)
	fmt.Println("Returned change: ", change+1)
	return change + 1
}
