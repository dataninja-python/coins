

/*
import (
	"sort"
)

// sums all integers in array
func sumArray(arrayToSum []int) int {
	theSum := 0
	if len(arrayToSum) == 0 {
		panic("Array sent to sumArray is empty")
	}
	for _, v := range arrayToSum {
		theSum += v
	}
	return theSum
}

func createPermutations(coinsArray []int) (permuts [][]int) {
	// create variable to store permutations
	// permutationStorage := [][]int{}
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(coinsArray); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(coinsArray, 0)
	return permuts
}

func findSum(currentChange int, coinArray []int, coinArrHolder [][]int) bool {
	// initialize a variable to let us know if we found the
	// currentChange amount
	found := false
	// arrayHolder := [][]int{}
	workingArray := coinArray
	// check to see if we can solve it with a single function
	checkSingle := func(change int, arr []int) bool {
		for _, v := range arr {
			if v > change {
				return false
			}
			if v == change {
				return true
			}
		}
		return false
	}
	// if true then currentchange == 1 of the numbers in the array
	if checkSingle(currentChange, workingArray) {
		return true
	} else {
		for _, outer := range coinArrHolder {
			theSum := 0
			for _, inner := range outer {
				theSum += inner
				switch {
				case currentChange == theSum:
					return true
				case theSum > currentChange:
					continue
				}
			}
		}
	}
	return found
}

func NonConstructibleChange(coins []int) int {
	// create minimum change variable
	// initialize it with 1
	minChange := 1
	// count the number of coins in the array
	numberOfCoins := len(coins)
	// if no coins provided return 1
	if numberOfCoins == 0 {
		return 1
	}
	// sort coins to ensure the permutations we perform are in order
	// fmt.Println("Unsorted coins array: ", coins)
	sort.Ints(coins)
	// fmt.Println("Sorted coins array: ", coins)
	// maximum change possible to create by adding all coins in array
	maxChange := sumArray(coins)
	// fmt.Println("Max possible change: ", maxChange)

	canMake := true

	arrayHolder := createPermutations(coins)
	// fmt.Println(arrayHolder)
	for i := 1; i <= maxChange; i++ {
		// count from 1 through the max change possible to create
		// fmt.Println(i)
		// check if change can be created
		// if it can return the current i as minChange
		canMake = findSum(i, coins, arrayHolder)
		if !canMake {
			minChange = i
			return minChange
		}
		// if we reach the last max change number then the minimum change
		// we can't calculate is + 1 the max change number
		if i == maxChange {
			minChange = maxChange + 1
		}
	}
	return minChange
}
*/
/*
func NonConstructibleChange2(coins []int) int {
	minChange := 1
	// numberOfCoins := len(coins)

	/*sumCoins := func(a []int) int {
		theSum := 0
		for _, v := range a {
			theSum += v
		}
		return theSum
	}*/

// maxCoins := sumCoins(coins)
// fmt.Println(minChange)
// fmt.Println(numberOfCoins)
// fmt.Println(maxCoins)
// if we reach the last max change number then the minimum change
// we can't calculate is + 1 the max change number
// if i == maxChange {
// 	minChange = maxChange + 1
// }
// return minChange
// }
