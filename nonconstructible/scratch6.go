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

func TestArray6(arr []int) int {
	currentChange := 0
	// continueCount := true
	fmt.Println("Starting array: ", arr)
	// sort array
	sort.Ints(arr)
	fmt.Println("Sorted array: ", arr)
	// check against base case scenarios that
	// make the answer automatically 1
	return currentChange
}
