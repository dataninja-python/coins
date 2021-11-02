
/*type comparison struct {
	testNum    int
	largestSum int
	startInd   int
	testArr    []int
}
func permut(t []int, p [][]int) [][]int {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			p = append(p, append([]int{}, a...))
		} else {
			for i := k; i < len(t); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(t, 0)
	return p
}

func canCreateSingle(testNum int, numArr []int) bool {
	for _, v := range numArr {
		if testNum == v {
			return true
		}
	}
	return false
}

func WalkArray(arr []int) bool {
	sort.Ints(arr)
	fmt.Println("Sorted array: ", arr)
	currentChange := 0
	continueCount := true
	// tempArr := []int{}

	for continueCount {
		// increment to create first change for comparison
		currentChange++
		var permutations [][]int
		// can create current number from a single loop through the array
		// continueCount = canCreateSingle(currentChange, arr)
		permutations = permut(arr, permutations)
		for _, a := range permutations {
			theSum := 0
			for _, v := range a {
				theSum = v
				fmt.Println(theSum)
				if currentChange == theSum {
					continue
				}
			}
		}
		continueCount = false
	}
	fmt.Println("Change can't create: ", currentChange)
	return false
}
*/
