package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard, within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, ok := cb[file]
	if !ok {
		return 0
	}
	count := 0
	for i := 0; i < len(f); i++ {
		if f[i] {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard, within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	count := 0
	index := rank - 1 // ranks are 1–8, slice index is 0–7

	for _, file := range cb {
		if index < len(file) && file[index] {
			count++
		}
	}
	return count
}


// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for range cb {
		count += 8
	}
	return count
}


// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total := 0
	for file := range cb {
		total += CountInFile(cb, file)
	}
	return total
}


