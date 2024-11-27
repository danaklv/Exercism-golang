package tasks

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool 

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string] File 

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count := 0
    for k, v := range cb {
        if k == file {
            for _, r := range v {
				if r {
					count ++
				}
			}
        }
    }
    return count
}


// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    count := 0
    for _, v := range cb {
            for i, r := range v {
				if i + 1 == rank  && r {
					count ++
				}
			}
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    	count := 0
    for _, v := range cb {
            for _, r := range v {
				if r {
                    count ++
                }
			}
    }
    return count	
}