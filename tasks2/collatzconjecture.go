package tasks2

import "fmt"

func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, fmt.Errorf("n should be possitive number")
	} else if n == 1 {
		return 0, nil
	}
	steps := 0

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps, nil
}
