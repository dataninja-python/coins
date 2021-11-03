

// type compare struct {
// 	startInd    int
// 	nextInd     int
// 	lastInd     int
// 	currentNum  int
// 	originalArr []int
// 	currentArr  []int
// }

/*func removeIndex(aSlice []int, index int) []int {
	temp := aSlice
	return append(temp[:index], temp[index+1:]...)
}*/

/*func TestArray(arr []int) int {
fmt.Println("Original array: ", arr)
sort.Ints(arr)
fmt.Println("Sorted array: ", arr)
currentChange := 0
continueCount := true
workingArr := make([]int, len(arr))
copy(workingArr, arr)
if len(arr) == 0 {
	return 1
}
if workingArr[0] != 1 {
	return 1
}
theSum := 0
fmt.Println("Starting Sum: ", theSum)
fmt.Println("Working array: ", workingArr)
for continueCount {
	// increment to create first change for comparison
	currentChange++
	for i := 0; i < len(arr); i++ {
		start := i
		theSum = 0
		compare := workingArr[start]
		temp := removeIndex(workingArr, start)
		// fmt.Println(i, start, next, compare, temp)
		// do comparison
		theSum += compare
		fmt.Println("Initial sum: ", theSum)
		if currentChange == compare {
			continue
		}
		for _, v := range temp {
			theSum += v
			fmt.Println("Current sum: ", theSum)
			if theSum == currentChange {
				continue
			}
			if theSum > currentChange {
				continue
			}
		}
		// reset for next comparison
		copy(workingArr, arr)
		// fmt.Println("reset: ", workingArr)
	}
	continueCount = false
	/*if currentChange == 1 {
		continueCount = false
	}*/
// }
// return currentChange
// }
