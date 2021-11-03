

/*type makeChange struct {
	originalSlice []int
	maxSwaps      int
	currentSwaps  int
	change        int
	currentSlice  []int
	sumsSlice     []int
	changeMap     map[int]bool
	canMakeChange bool
}

func removeIndex(aSlice []int, start int, stop int) []int {
	// stop non inclusive
	if start+stop > len(aSlice) {
		return append(aSlice[:start], aSlice[len(aSlice):]...)
	}
	return append(aSlice[:start], aSlice[start+stop:]...)
}*/

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

/*func createSumArray(a []int) map[int]bool {
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
}*/

/*func makeFirstLast(aHolder makeChange) makeChange {
	// perform swap
	// tempSlice := make([]int, len(aHolder.originalSlice))
	// copy(tempSlice, aHolder.originalSlice)
	// tempSlice2 := make([]int, len(aHolder.currentSlice))
	// copy(tempSlice2, aHolder.currentSlice)
	firstItem := aHolder.currentSlice[:1]
	// fmt.Println(firstItem)
	// fmt.Println(aHolder.currentSlice)
	aHolder.currentSlice = append(aHolder.currentSlice[:0],
		aHolder.currentSlice[1:]...)
	// fmt.Println(aHolder.currentSlice)
	aHolder.currentSlice = append(aHolder.currentSlice, firstItem...)
	// fmt.Println(aHolder.currentSlice)
	// add +1 to swaps
	aHolder.currentSwaps++
	// if swaps == max swaps and still haven't created the change then change can make change to false and end loop
	return aHolder
}*/

/*func getPermutations(cHolder makeChange) {
	var sumExist = make(map[int]bool)
	sumSlice := []int{}
	cHolder.sumsSlice = sumSlice
	cHolder.changeMap = sumExist
	// fmt.Println(cHolder)
	// loop through swapping first to last digit until can't do it anymore
	for cHolder.canMakeChange {
		// increase change + 1
		cHolder.change++
		// see if can make sum using current arrangement
		// loop from 1 through < len(arr) to capture the gap of each array
		for g := 1; g < len(cHolder.originalSlice); g++ {
			fmt.Println("current gap from start: ", g)
			fmt.Println("current slice working on: ", cHolder.currentSlice)
			// capture temp to be able to reset
			// tempSlice := make([]int, len(cHolder.currentSlice))
			// copy(tempSlice, cHolder.currentSlice)
			// starting from beginning of slice, loop through all
			// viable permutations
			for i := 0; i < len(cHolder.originalSlice); i++ {
				next := i + g
				if next > len(cHolder.originalSlice) {
					continue
				} else {
					// theSum := 0
					startSlice := cHolder.currentSlice[i:next]
					sumSlice := cHolder.currentSlice[next:]
					fmt.Println("Index: ", i, "Next: ", next,
						"StartSlice: ", startSlice, "SumSlice: ",
						sumSlice, "CurrentSlice: ", cHolder.currentSlice)
				}
			}
			// make all slices and sum slices to check if in map as index
			// if not in map add it, else skip
		}

		// check if have reached maximum swaps == back to original slice arrangement == can't make change
		if cHolder.currentSwaps == cHolder.maxSwaps {
			fmt.Println(cHolder)
			cHolder.canMakeChange = false
		}
		// if can't make change with current configuration switch last and first digits
		cHolder = makeFirstLast(cHolder)
		// fmt.Println(cHolder)
	}
}*/

/*
func TestArray2(arr []int) int {
	currentChange := 0
	// continueCount := true
	// fmt.Println("Original array: ", arr)
	sort.Ints(arr)
	fmt.Println("Sorted array: ", arr)
	// check against base case scenarios
	switch {
	case len(arr) == 0 || arr[0] != 1:
		return 1
	default:
		tempSlice := make([]int, len(arr))
		copy(tempSlice, arr)
		changeHolder := makeChange{originalSlice: tempSlice, maxSwaps: len(arr), currentSwaps: 0, change: currentChange, currentSlice: arr, canMakeChange: true}
		// run a 2D array that runs through all number permutations
		getPermutations(changeHolder)
	}

	// create a spread that counts from 1 to 6 the max spread there can be
	/*for spread := 1; spread < len(arr); {
		fmt.Println("spread is: ", spread)
		aTemp := removeIndex(arr, 0, spread)
		fmt.Println(aTemp)
		spread++
	}*/
// comparisonMap := createSumArray(arr)
// fmt.Println(comparisonMap)
// fmt.Println("Comparison array: ", comparisonArray)
// sort.Ints(comparisonArray)
// fmt.Println("Sorted array: ", comparisonArray)
// comparisonArray = removeDupes(comparisonArray)
// fmt.Println("Deduplicated: ", comparisonArray)
/*for continueCount {
	// increment to create first change for comparison
	currentChange++

}*/
/*return currentChange
}*/
