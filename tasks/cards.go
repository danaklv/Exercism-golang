package tasks

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	for i, r := range slice {
		if i == index {
			return r
		}
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	for i, _ := range slice {
		if i == index {
			slice[i] = value
			return slice
		}
	}
	slice = append(slice, value)
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var sliceValues []int
	sliceValues = append(sliceValues, values...)
	sliceValues = append(sliceValues, slice...)
	return sliceValues
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice)-1 {
		return slice
	}
	copy(slice[index:], slice[index+1:])

	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	return slice

}
