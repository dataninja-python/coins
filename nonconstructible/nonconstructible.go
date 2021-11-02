package nonconstructible

import (
	"fmt"
	"sort"
)

type workVars struct {
	changeMakeable     bool
	swaps              int
	workSlice          []int
	sumMap             map[int]bool
	sumSlice           []int
	originalSlice      []int
	lastChangeMakeable int
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

func getCutSlice(aSlice []int, startInc int, stopExc int) []int {
	temp := []int{}
	// skip part to be removed
	for index, value := range aSlice {
		if index >= startInc && index < stopExc {
			continue
		} else {
			temp = append(temp, value)
		}
	}
	return temp
}

func getMaxChangeMakeable(theWorker workVars) workVars {
	sort.Ints(theWorker.sumSlice)
	theWorker.changeMakeable = true
	counter := 0
	for i := 0; i < len(theWorker.sumSlice); i++ {
		counter++
		if counter == theWorker.sumSlice[i] &&
			counter > theWorker.lastChangeMakeable {
			fmt.Println("Adding ", counter, " to lastChangeMakeable")
			theWorker.lastChangeMakeable = counter
		}

		if counter > theWorker.sumSlice[i] {
			if _, ok := theWorker.sumMap[counter]; !ok {
				fmt.Println(counter, "can't be created with available elements")
				return theWorker
			}
		}
	}
	return theWorker
}

func getSumsPerSlice(startInc, stopExc int, aWorker workVars) workVars {
	// create a new temporary slice for this function
	tempSlice := getNewSlice(aWorker.workSlice)
	// slice to sum before going through loop of each item
	baseSumSlice := aWorker.workSlice[startInc:stopExc]
	// add base sum to each item in this slice
	// NOTE: this is the slice that does not contained items added to baseSum
	remainingSlice := getCutSlice(tempSlice, startInc, stopExc)

	baseSum := 0
	for _, value := range baseSumSlice {
		baseSum += value
		if _, ok := aWorker.sumMap[baseSum]; !ok {
			aWorker.sumMap[baseSum] = true
			aWorker.sumSlice = append(aWorker.sumSlice, baseSum)
		}
	}
	for i := 0; i < len(remainingSlice); i++ {
		baseSum += remainingSlice[i]
		if _, ok := aWorker.sumMap[baseSum]; !ok {
			aWorker.sumMap[baseSum] = true
			aWorker.sumSlice = append(aWorker.sumSlice, baseSum)
		}
		baseSum -= remainingSlice[i]
	}
	aWorker = getMaxChangeMakeable(aWorker)
	return aWorker
}

func getUniqueMakeableChange(begIndex int, begSpread int,
	aWorker workVars) workVars {
	continueLoop := true
	for continueLoop {
		for spread := begSpread; spread <= len(aWorker.originalSlice); spread++ {
			for sliceIndex := begIndex; sliceIndex < len(aWorker.originalSlice); sliceIndex++ {
				// create variables to hold looping through each combination of
				// the sorted slice to get all sums
				startSliceInclusive := sliceIndex
				// must protect against going outside the bounds of slice exclusive
				// this should be len(slice)
				endSliceExclusive := min(startSliceInclusive+spread, len(aWorker.originalSlice))
				aWorker = getSumsPerSlice(startSliceInclusive,
					endSliceExclusive, aWorker)
			}
		}
		continueLoop = false
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
	worker := workVars{changeMakeable: true, swaps: 0, workSlice: s1, sumMap: sumTracker, originalSlice: s2, lastChangeMakeable: minChange}
	fmt.Println("Initial worker workVars: ", worker, "\n")
	// start finding all unique change combinations using an inner loop
	// to loop through eah index in a coin and a spread to extract the
	// portion of the coin array to count
	worker = getUniqueMakeableChange(0, 1, worker)
	fmt.Println("End worker workVars: ", worker, "\n")
	return minChange
}
