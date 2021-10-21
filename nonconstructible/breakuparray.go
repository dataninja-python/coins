package nonconstructible

import (
	"fmt"
	"sort"
)

func WalkArray(arr []int) {
	sort.Ints(arr)
	temp := []int{}
	// tempMap := map[int][][]int{}
	// row := 1
	// you know you are going to do something for the length of the array
	for i, v := range arr {
		fmt.Println("The current index: ", i)
		fmt.Println("The current value: ", v)
		if i+1 == len(arr)+1 {
			continue
		} else {
			temp = arr[i+1:]
		}
		fmt.Println("The current short array: ", temp)
	}
}
