package tasks2

func Score(word string) int {
	score := 0
	for _, w := range word {
		score += Rate(w)

	}
	return score

}

func Rate(r rune) int {

	 if  r >= 97 && r <=122 {
		r -= 32
	 }
	

	switch r {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
