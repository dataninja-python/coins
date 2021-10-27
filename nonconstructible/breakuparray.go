package nonconstructible

import (
	"fmt"
	"sort"
)

/*type comparison struct {
	testNum    int
	largestSum int
	startInd   int
	testArr    []int
}*/

func canCreateChange(testNum int, numArr []int) bool {
	theSum := 0
	for _, v := range numArr {
		theSum += v
		// if we can create the change return true
		if testNum == theSum {
			return true
		}
		// if the sum of numbers in array is greater than the testNum
		// we can not create the change
		if theSum > testNum {
			return false
		}
	}
	return false
}

func WalkArray(arr []int) bool {
	sort.Ints(arr)
	fmt.Println("Sorted array: ", arr)
	currentChange := 0
	continueCount := true

	for continueCount {
		// increment to create first change for comparison
		currentChange++
		if currentChange == 10 {
			continueCount = false
		}
	}
	fmt.Println("Change can't create: ", currentChange)
	return false
}
