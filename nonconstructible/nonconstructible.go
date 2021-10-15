package nonconstructible

import (
	"fmt"
	"sort"
)

func NonConstructibleChange(coins []int) int {
	minChange := -1
	sort.Ints(coins)
	fmt.Println(minChange)
	total := sumAll(coins)
	fmt.Println(total)
	return coins[0]
}

func sumAll(coins []int) int {
	aSum := 0
	for _, v := range coins {
		aSum += v
	}
	return aSum
}
