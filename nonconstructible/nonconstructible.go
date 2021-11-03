package nonconstructible

import (
	"fmt"
	"sort"
)

func getNewSlice(originalSlice []int) []int {
	tempSlice := make([]int, len(originalSlice))
	copy(tempSlice, originalSlice)
	return tempSlice
}

func NonConstructibleChange(coins []int) int {
	// initialize key values
	minChange := 0
	startingSlice := getNewSlice(coins)

	// sort array
	sort.Ints(startingSlice)

	// if array is empty or first element != 1
	// then you can't construct 1
	if len(coins) == 0 || coins[0] != 1 {
		return 1
	}

	fmt.Println(startingSlice, coins)
	return minChange
}
