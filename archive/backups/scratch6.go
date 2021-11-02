

/*type workVars struct {
	changeMakeable bool
	swaps          int
	workSlice      []int
	sumMap         map[int]bool
	sumSlice       []int
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

		if _, ok := theWorker.sumMap[counter]; !ok {
			return theWorker
		}
	}
	return theWorker
}

func getNewSlice(originalSlice []int) []int {
	tempSlice := make([]int, len(originalSlice))
	copy(tempSlice, originalSlice)
	return tempSlice
}*/

/*func removeFromSlice(aSlice []int, startInclusive int,
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
}*/

/*
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
}*/

/*
func minNonChange(start int, aWorker workVars) int {
	lastMatch := 0
	continueLoop := true
	for continueLoop {
		start++
		for i, v := range aWorker.sumSlice {
			if start == v {
				lastMatch = start
				break
			} else if i >= len(aWorker.sumSlice)-1 {
				// fmt.Println(lastMatch)
				continueLoop = false
			}
		}
	}
	return lastMatch + 1
}

func TestArray6(arr []int) int {
	currentChange := 0
	// continueCount := true
	// fmt.Println("Starting array: ", arr)
	// sort array
	sort.Ints(arr)
	// fmt.Println("Sorted array: ", arr)
	// check against base case scenarios that
	// make the answer automatically 1
	if len(arr) == 0 || arr[0] != 1 {
		return 1
	}

	// above cases essentially ensure that change we want to make
	// must at least be 1
	currentChange = 1

	s1 := getNewSlice(arr)
	sumTracker := make(map[int]bool)
	worker := workVars{changeMakeable: true, swaps: 0, workSlice: s1, sumMap: sumTracker}
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
			// fmt.Println(currentChange, aSum)
		} else {
			worker.changeMakeable = false
			// fmt.Println(currentChange, aSum)
		}
	}

	if worker.swaps <= len(arr) {
		worker.workSlice = moveFirstElementToEnd(worker.workSlice)
		worker.swaps++
		worker.changeMakeable = true
	} else {
		worker.changeMakeable = false
	}
	// fmt.Println(worker.sumSlice)
	currentChange = minNonChange(currentChange, worker)
	return currentChange
}
*/
