// Package queenattack checks whether two queens can attack each other
package queenattack

import (
	"errors"
	"strconv"
)

// pos is a position on a chess board
type pos struct {
	row int
	col int
}

//the length we expected from a sigil
const expectedLen = 2

//the test version
const testVersion = 1

// convertToPos takes a string and gives back a position
func convertToPos(sigil string) (pos, error) {
	if len(sigil) != expectedLen {
		return pos{}, errors.New("Invalid position on chess board")
	}

	row := int(sigil[0] - 'a' + 1)
	col, err := strconv.Atoi(sigil[1:])

	// If the conversion fails, return an error
	if err != nil {
		return pos{}, err
	}

	// Bounds check the row
	if row < 1 || row > 8 {
		return pos{}, errors.New("Out of bounds position")
	}

	// Bounds check the column
	if col < 1 || col > 8 {
		return pos{}, errors.New("Out of bounds positions")
	}

	return pos{row, col}, nil
}

// CanQueenAttack checks whether a pair of queens are in attacking positions
func CanQueenAttack(q1, q2 string) (bool, error) {
	p1, err := convertToPos(q1)

	// bubble up the error if we had one
	if err != nil {
		return false, err
	}

	p2, err := convertToPos(q2)

	// again, bubble up the error if we had one
	if err != nil {
		return false, err
	}

	// Check if the queens are at the same place
	if p1 == p2 {
		return false, errors.New("Queens on the same position")
	}

	// Check if they are in the same column or row
	if p1.row == p2.row || p1.col == p2.col {
		return true, nil
	}

	rowDiff := p1.row - p2.row
	colDiff := p1.col - p2.col

	//if these are the same, or just negatives of each other, then they are on an attacking diagonal
	if rowDiff == colDiff || rowDiff == -colDiff {
		return true, nil
	}

	return false, nil

}
