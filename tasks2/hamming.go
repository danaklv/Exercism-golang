package tasks2

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("lengths sbould be equal")
	}
	count, i := 0, 0

	for i < len(a) {
		if a[i] != b[i] {
			count++
		}
		i++
	}
	return count, nil
}
