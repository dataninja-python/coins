package nonconstructible

import (
	"fmt"
	"sort"
)

func permutation(coinsArr []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(coinsArr); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(coinsArr, 0)

	return permuts
}

func NonConstructibleChange(coins []int) int {
	// create minimum change variable
	// initialize it with 1
	minChange := 1
	// if no coins provided return 1
	if len(coins) == 0 {
		return 1
	}
	sort.Ints(coins)
	// process elements from left to right
	noOfPerms := permutation(coins)
	fmt.Println("Number of Permutation:", len(noOfPerms))
	for i := 0; i < len(noOfPerms); i++ {
		fmt.Println(noOfPerms[i])
	}
	return minChange
}
