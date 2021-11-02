package nonconstructible

import (
	"fmt"
	"sort"
)

type workVars struct {
	changeMakeable bool
	swaps          int
	workSlice      []int
	sumMap         map[int]bool
	sumSlice       []int
	originalSlice  []int
}

func getNewSlice(originalSlice []int) []int {
	tempSlice := make([]int, len(originalSlice))
	copy(tempSlice, originalSlice)
	return tempSlice
}

// source: https://stackoverflow.com/questions/27516387/what-is-the-correct-way-to-find-the-min-between-two-integers-in-go
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getSumsPerSlice(startInc, stopExc int, aWorker workVars) workVars {
	return aWorker
}

func getUniqueMakeableChange(begIndex int, begSpread int,
	aWorker workVars) workVars {
	for spread := begSpread; spread < len(aWorker.originalSlice); spread++ {
		for sliceIndex := begIndex; sliceIndex < len(aWorker.originalSlice); sliceIndex++ {
			// create variables to hold looping through each combination of
			// the sorted slice to get all sums
			startSliceInclusive := sliceIndex
			// must protect against going outside the bounds of slice exclusive
			// this should be len(slice)
			endSliceExclusive := min(startSliceInclusive+spread, len(aWorker.originalSlice))
			fmt.Println("start include: ", startSliceInclusive,
				"stop exclude: ", endSliceExclusive)
		}
	}
	return aWorker
}

func NonConstructibleChange(coins []int) int {
	// set starting number to 0
	minChange := 0
	// sort array
	sort.Ints(coins)

	// if array is empty or first element != 1
	// then you can't construct 1
	if len(coins) == 0 || coins[0] != 1 {
		return 1
	}

	// create workVars type to store key data
	s1 := getNewSlice(coins)
	s2 := getNewSlice(coins)
	sumTracker := make(map[int]bool)
	worker := workVars{changeMakeable: true, swaps: 0, workSlice: s1, sumMap: sumTracker, originalSlice: s2}
	fmt.Println("Initial worker workVars: ", worker)
	// start finding all unique change combinations using an inner loop
	// to loop through eah index in a coin and a spread to extract the
	// portion of the coin array to count
	worker = getUniqueMakeableChange(0, 1, worker)
	return minChange
}
