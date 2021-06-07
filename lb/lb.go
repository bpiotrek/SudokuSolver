package lb

type CellType = string
type FlatBoardType = [81]CellType
type BoardType = [9][9]CellType

func GetBox(i, j int) int {
	return i/3*3 + j/3
}

func GetEmpty(board *BoardType, row, col int) (int, int) {
	for row < 9 {
		if (*board)[row][col] == "." {
			break
		}
		row, col = getNextPos(row, col)
	}
	return row, col
}

func getNextPos(row, col int) (int, int) {
  col++
	row = row + col/9
	col = col % 9
	return row, col
}
