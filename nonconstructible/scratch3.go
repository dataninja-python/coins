package nonconstructible

import (
	"fmt"
	"sort"
)

// type compare struct {
// 	startInd    int
// 	nextInd     int
// 	lastInd     int
// 	currentNum  int
// 	originalArr []int
// 	currentArr  []int
// }

func removeIndex(aSlice []int, start int, stop int) []int {
	// stop non inclusive
	if start+stop > len(aSlice) {
		return append(aSlice[:start], aSlice[len(aSlice):]...)
	}
	return append(aSlice[:start], aSlice[start+stop:]...)
}

/*func removeDupes(intSlice []int) []int {
	keys := make(map[int]bool)
	deduped := []int{}
	// source: geeksforgeeks.org
	// if the key(values of the slice) is not equal
	// to the value in new slice then we append it,
	// else skip it
	for _, aInt := range intSlice {
		if _, ok := keys[aInt]; !ok {
			// fmt.Println("Unique value: ", aInt)
			keys[aInt] = true
			deduped = append(deduped, aInt)
		}
	}
	return deduped
}*/

func createSumArray(a []int) map[int]bool {
	output := make(map[int]bool)
	workingArr := make([]int, len(a))
	copy(workingArr, a)
	theSum := 0
	fmt.Println("Starting Sum: ", theSum)
	fmt.Println("Working array: ", workingArr)
	var finalSlice []int
	for s := 1; s < len(a); s++ {
		for i := 0; i < len(a); i++ {
			start := i
			theSum = 0
			compare := workingArr[start]
			fmt.Println("Working array: ", workingArr)
			// append compare
			if _, ok := output[compare]; !ok {
				output[compare] = true
				finalSlice = append(finalSlice, compare)
			}
			fmt.Println("Final slice with compare: ", finalSlice)
			temp := removeIndex(workingArr, start, s)
			theSum += compare
			for _, v := range temp {
				theSum += v
				if _, ok := output[theSum]; !ok {
					output[theSum] = true
					finalSlice = append(finalSlice, theSum)
				}
				fmt.Println("Final slice + next sum: ", finalSlice)
			}
			// reset for next comparison
			copy(workingArr, a)
			fmt.Println("reset: ", workingArr)
		}
	}
	sort.Ints(finalSlice)
	// remove duplicates
	fmt.Println("Final unsorted array: ", finalSlice)
	return output
}

func TestArray2(arr []int) int {
	currentChange := 0
	// continueCount := true
	fmt.Println("Original array: ", arr)
	sort.Ints(arr)
	fmt.Println("Sorted array: ", arr)
	if len(arr) == 0 {
		return 1
	}
	if arr[0] != 1 {
		return 1
	}
	// create a spread that counts from 1 to 6 the max spread there can be
	/*for spread := 1; spread < len(arr); {
		fmt.Println("spread is: ", spread)
		aTemp := removeIndex(arr, 0, spread)
		fmt.Println(aTemp)
		spread++
	}*/
	comparisonMap := createSumArray(arr)
	fmt.Println(comparisonMap)
	// fmt.Println("Comparison array: ", comparisonArray)
	// sort.Ints(comparisonArray)
	// fmt.Println("Sorted array: ", comparisonArray)
	// comparisonArray = removeDupes(comparisonArray)
	// fmt.Println("Deduplicated: ", comparisonArray)
	/*for continueCount {
		// increment to create first change for comparison
		currentChange++

	}*/
	return currentChange
}
