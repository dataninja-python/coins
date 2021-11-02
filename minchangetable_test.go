package main

import (
	"change/nonconstructible"
	"testing"
)

// source: https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package
type changeTest struct {
	slice    []int
	expected int
}

var changeTests = []changeTest{
	changeTest{[]int{1, 2, 5}, 4},
	changeTest{[]int{}, 1},
	changeTest{[]int{87}, 1},
	changeTest{[]int{5, 7, 1, 1, 2, 3, 22}, 20},
	changeTest{[]int{1, 1, 1, 1, 1}, 5},
	changeTest{[]int{1, 5, 1, 1, 1, 10, 15, 20, 100}, 45},
	changeTest{[]int{6, 4, 5, 1, 1, 8, 9}, 3},
	changeTest{[]int{5, 6, 1, 1, 2, 3, 4, 9}, 32},
	changeTest{[]int{5, 6, 1, 1, 2, 3, 43}, 19},
	changeTest{[]int{1, 1}, 3},
	changeTest{[]int{2}, 1},
	changeTest{[]int{1}, 2},
	changeTest{[]int{109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4}, 87},
	changeTest{[]int{1, 2, 3, 4}, 11},
}

func TestChange(t *testing.T) {
	for _, test := range changeTests {
		if output := nonconstructible.NonConstructibleChange(test.slice); output != test.expected {
			t.Errorf("Output %d not equal to expected %d",
				output, test.expected)
		}
	}
}
