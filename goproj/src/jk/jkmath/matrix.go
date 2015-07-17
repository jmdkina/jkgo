package jkmath

// package main

import (
	"errors"
	"fmt"
	"jk/jklog"
)

type Matrix struct {
	// The first is rows, then cols.
	Items [][]int
}

/*
 * items first colume then row, from left to right, from top to bottom.
 */
func NewMatrix(cols, rows int, items []int) *Matrix {
	matrix := Matrix{}
	matrix.Items = make([][]int, rows)
	// Fill the matrix
	for i := 0; i < rows; i++ {

		for j := 0; j < cols; j++ {
			matrix.Items[i] = append(matrix.Items[i], items[i*cols+j])
		}
	}
	return &matrix
}

func (m *Matrix) Display() {
	jklog.L().Infof("The max of cols, rows [%d,%d]\n", len(m.Items[0]), len(m.Items))
	for i := 0; i < len(m.Items); i++ {
		for j := 0; j < len(m.Items[i]); j++ {
			fmt.Printf("     %d", m.Items[i][j])
		}
		fmt.Println("")
	}
}

// The size of matrix with []int , cols and rows
func (m *Matrix) Size() []int {
	size := []int{
		len(m.Items[0]),
		len(m.Items),
	}
	return size
}

/*
 * row and col is the index from 1 to ...
 */
func (m *Matrix) Item(row, col int) (int, error) {
	irow := row - 1
	icol := col - 1
	if irow < 0 || irow >= len(m.Items) || icol < 0 || icol >= len(m.Items[0]) {
		return 0, errors.New("Wrong index")
	}
	return m.Items[irow][icol], nil
}

/*
func main() {
	matrix := []int{1, 0, -1, 2, 2, -3, 3, 4, 0}
	theM := NewMatrix(3, 3, matrix)
	theM.Display()

	jklog.L().Infoln("\n")

	nM := NewMatrix(4, 3, []int{1, 0, 4, 2, 1, 1, 3, 1, 0, 0, 2, 2})
	nM.Display()

	jklog.L().Infof("The matrix size is [%d,%d]\n", nM.Size()[0], nM.Size()[1])
	v, err := nM.Item(4, 1)
	if err != nil {
		jklog.L().Errorln("wrong item as : ", err)
	} else {
		jklog.L().Infof("The matrix item of [2,1] is [%d]\n", v)
	}
}
*/
