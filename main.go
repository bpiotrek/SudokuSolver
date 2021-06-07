package main

import (
	"fmt"
	"main/lb"
	"strconv"
	"strings"

	"reflect"

	"github.com/bits-and-blooms/bitset"
)

type cellType = lb.CellType
type flatBoardType = lb.FlatBoardType
type boardType = lb.BoardType

func main() {
	flat_board := flatBoardType{
		"5", "3", ".", ".", "7", ".", ".", ".", ".",
		"6", ".", ".", "1", "9", "5", ".", ".", ".",
		".", "9", "8", ".", ".", ".", ".", "6", ".",
		"8", ".", ".", ".", "6", ".", ".", ".", "3",
		"4", ".", ".", "8", ".", "3", ".", ".", "1",
		"7", ".", ".", ".", "2", ".", ".", ".", "6",
		".", "6", ".", ".", ".", ".", "2", "8", ".",
		".", ".", ".", "4", "1", "9", ".", ".", "5",
		".", ".", ".", ".", "8", ".", ".", "7", "9",
	}
	flat_expected := flatBoardType{
		"5", "3", "4", "6", "7", "8", "9", "1", "2",
		"6", "7", "2", "1", "9", "5", "3", "4", "8",
		"1", "9", "8", "3", "4", "2", "5", "6", "7",
		"8", "5", "9", "7", "6", "1", "4", "2", "3",
		"4", "2", "6", "8", "5", "3", "7", "9", "1",
		"7", "1", "3", "9", "2", "4", "8", "5", "6",
		"9", "6", "1", "5", "3", "7", "2", "8", "4",
		"2", "8", "7", "4", "1", "9", "6", "3", "5",
		"3", "4", "5", "2", "8", "6", "1", "7", "9",
	}
	board := splice(&flat_board)
	expected := splice(&flat_expected)
	fmt.Printf("initial:\n%sexpected:\n%s", printBoard(board), printBoard(expected))
	SudokuSolve(board)
	res := "different"
	if reflect.DeepEqual(board, expected) {
		res = "identical"
	}
	fmt.Printf("actual:\n%s\nboards are %s\n", printBoard(board), res)
}

func printBoard(b *boardType) string {
	var sb strings.Builder
	for _, row := range *b {
		for _, cell := range row {
			fmt.Fprintf(&sb, "%s, ", cell)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func splice(b *flatBoardType) *boardType {
	var r boardType
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			r[i][j] = (*b)[i*9+j]
		}
	}
	return &r
}

func SudokuSolve(board *boardType) {
	row_contains := make([]*bitset.BitSet, 9)
	col_contains := make([]*bitset.BitSet, 9)
	box_contains := make([]*bitset.BitSet, 9)

	for i := 0; i < len(row_contains); i++ {
		row_contains[i] = bitset.New(9)
		col_contains[i] = bitset.New(9)
		box_contains[i] = bitset.New(9)
	}

	for i, row := range *board {
		for j, cell := range row {
			if cell != "." {
				digit_, _ := strconv.Atoi(cell)
				digit := uint(digit_) - 1
				row_contains[i].Set(digit)
				col_contains[j].Set(digit)
				box_contains[lb.GetBox(i, j)].Set(digit)
			}
		}
	}

	solve(board, 0, 0, &row_contains, &col_contains, &box_contains)
}

func solve(board *boardType, row_start, col_start int, row_contains, col_contains, box_contains *[]*bitset.BitSet) bool {
	row, col := lb.GetEmpty(board, row_start, col_start)
	if row > 8 {
		return true
	}

	box := lb.GetBox(row, col)
	contains := (*row_contains)[row].Union((*col_contains)[col])
	contains.InPlaceUnion((*box_contains)[box])
	if contains.All() {
		return false
	}

	for i := 0; i < 9; i++ {
		j := uint(i)
		if !contains.Test(j) {
			(*board)[row][col] = strconv.Itoa(i + 1)
			(*row_contains)[row].Set(j)
			(*col_contains)[col].Set(j)
			(*box_contains)[box].Set(j)
			if solve(board, row, col, row_contains, col_contains, box_contains) {
				return true
			}
			(*row_contains)[row].SetTo(j, false)
			(*col_contains)[col].SetTo(j, false)
			(*box_contains)[box].SetTo(j, false)
		}
	}
	(*board)[row][col] = "."
	return false
}
