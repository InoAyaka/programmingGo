package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(1, 2, 5, 8, 23))
	fmt.Println(min(-1, 2, 5, -8, 23))
	fmt.Println(max())
	fmt.Println(min())

	fmt.Println(maxNeedArg(1, 2, 5, 8, 23))
	fmt.Println(minNeedArg(-1, 2, 5, -8, 23))
	fmt.Println(maxNeedArg(2))
	fmt.Println(minNeedArg(-100))
	//compile error
	//fmt.Println(maxNeedArg(2))
	//fmt.Println(minNeedArg())

}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("zero lenght slice")
	}

	var max int

	for _, val := range vals {
		if max < val {
			max = val
		}
	}

	return max, nil
}

func maxNeedArg(val int, vals ...int) int {
	max := val

	for _, val := range vals {
		if max < val {
			max = val
		}
	}

	return max
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("zero lenght slice")
	}

	var min int

	for _, val := range vals {
		if min > val {
			min = val
		}
	}

	return min, nil
}

func minNeedArg(val int, vals ...int) int {
	min := val

	for _, val := range vals {
		if min > val {
			min = val
		}
	}

	return min
}
