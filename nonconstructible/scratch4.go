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
}

func getNewSlice(originalSlice []int) []int {
	tempSlice := make([]int, len(originalSlice))
	copy(tempSlice, originalSlice)
	return tempSlice
}

func removeFromSlice(aSlice []int, startInclusive int,
	stopExclusive int) []int {
	temp := []int{}
	// skip part to be removed
	for index, value := range aSlice {
		if index >= startInclusive && index < stopExclusive {
			continue
		} else {
			temp = append(temp, value)
		}
	}
	return temp
}

func moveFirstElementToEnd(originalSlice []int) []int {
	tempSlice := getNewSlice(originalSlice)
	begSliceLength := len(tempSlice)
	// fmt.Println("Slice length @ start: ", begSliceLength)
	// get the first element from the original slice
	// Note: this is to avoid potetial memory issues
	// since slices are pointers to areas in memory
	var firstElement int
	if begSliceLength > 0 {
		firstElement = originalSlice[0]
	}
	// fmt.Println("First slice element: ", firstElement)
	// remove from start to end non-inclusive from slice
	begCut, endBefore := 0, 1
	tempSlice = removeFromSlice(tempSlice, begCut, endBefore)
	// fmt.Println("Interim slice: ", tempSlice,
	// "Original slice: ", originalSlice)
	tempSlice = append(tempSlice, firstElement)
	return tempSlice
}

func getSums(startInclusive int, endExclusive int, aWorker workVars) workVars {
	tempSlice := getNewSlice(aWorker.workSlice)
	// grab original sum to start
	startingSumSlice := aWorker.workSlice[startInclusive:endExclusive]
	addToSlice := removeFromSlice(tempSlice, startInclusive, endExclusive)
	// fmt.Println("initial slice to sum: ", startingSumSlice,
	// "slice to add each digit to: ", addToSlice,
	// "temp slice: ", tempSlice, "worker current slice: ", aWorker.workSlice)
	theSum := 0
	for _, v := range startingSumSlice {
		theSum += v
		if _, ok := aWorker.sumMap[theSum]; !ok {
			aWorker.sumMap[theSum] = true
			aWorker.sumSlice = append(aWorker.sumSlice, theSum)
			// fmt.Println("Appended sum: ", theSum)
		}
	}
	if len(addToSlice) > 0 {
		for i := 0; i < len(addToSlice); i++ {
			theSum += addToSlice[i]
			if _, ok := aWorker.sumMap[theSum]; !ok {
				aWorker.sumMap[theSum] = true
				aWorker.sumSlice = append(aWorker.sumSlice, theSum)
			}
			theSum -= addToSlice[i]
		}
	}
	return aWorker
}

func TestArray3(arr []int) int {
	currentChange := 0
	// continueCount := true
	fmt.Println("Starting array: ", arr)
	// sort array
	sort.Ints(arr)
	fmt.Println("Sorted array: ", arr)
	// check against base case scenarios that
	// make the answer automatically 1
	switch {
	case len(arr) == 0:
		// passed an empty array so should return 1
		currentChange = 1
		return currentChange
	case arr[0] != 1:
		// first item in sorted array is not equal to 1 so can't make 1 cent
		currentChange = 1
		return currentChange
	default:
		// start current change count from 2 because 1 cases taken care of
		currentChange = 1
		s1 := getNewSlice(arr)
		sumTracker := make(map[int]bool)
		// fmt.Println("original slice: ", arr, "new slice: ", s1)
		worker := workVars{changeMakeable: true, swaps: 0, workSlice: s1, sumMap: sumTracker}
		for worker.changeMakeable {
			currentChange++
			fmt.Println("Currently checking: ", currentChange)
			// run through each possible gap and cycle elements from front to back
			// Note: stop when back to original order
			for gap := 1; gap <= len(arr); gap++ {
				// take status at beginning of work
				// fmt.Println("Before swaps: ", "original slice: ", arr,
				// "new slice: ", worker.workSlice, "current swaps: ",
				// worker.swaps)
				// increment change to check if we can make change starting with 2
				for indexStart := 0; indexStart < len(arr); indexStart++ {
					// create all sums with current gap
					// begin by creating the starting index point
					begInclusive := indexStart
					// then, add the current gap or spread to have between
					// starting index element and last part of slice
					nextExclusive := begInclusive + gap
					if nextExclusive > 7 {
						nextExclusive = 7
					}
					// fmt.Println("beg = ", begInclusive, "spread = ", gap,
					// "nextExclusive = ", nextExclusive)
					// pass all to function that sums slice permutations
					worker = getSums(begInclusive, nextExclusive, worker)
				}
				// fmt.Println("After swaps: ", "original slice: ", arr,
				// "new slice: ", worker.workSlice, "current swaps: ",
				// worker.swaps)
			}
			sort.Ints(worker.sumSlice)
			// fmt.Println(worker.sumSlice)
			// check to see if current change is in map
			for _, aSum := range worker.sumSlice {
				if currentChange == aSum {
					worker.changeMakeable = true
					fmt.Println(currentChange, aSum)
				} else {
					worker.changeMakeable = false
					fmt.Println(currentChange, aSum)
				}
			}

			if worker.swaps <= len(arr) {
				worker.workSlice = moveFirstElementToEnd(worker.workSlice)
				worker.swaps++
				worker.changeMakeable = true
			}
			// fmt.Println(worker)
		}
	}
	return currentChange
}
