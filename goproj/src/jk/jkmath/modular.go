package jkmath

// package main

import (
	"errors"
	"fmt"
	"jk/jklog"
)

type Modular struct {
	n int
	b int
}

func New() *Modular {
	mod := Modular{}
	return &mod
}

func (m *Modular) printBaseExpansionResult(ret []int) {
	jklog.L().Infoln("The result is : ")
	rlen := len(ret)
	for i := rlen - 1; i >= 0; i-- {
		fmt.Printf("%d", ret[i])
		if i%4 == 0 {
			fmt.Printf(" ")
		}
	}

	fmt.Println()
}

func (m *Modular) BaseExpansion(n, b int) error {
	if b <= 1 || n <= 0 {
		return errors.New("base value is not valid.")
	}
	m.n = n
	m.b = b
	q := n
	var ret []int // save the result.
	for {
		// Mod last, exit
		if q == 0 {
			break
		}
		quo := q % b
		ret = append(ret, quo)
		q = q / b
	}
	// out. print the result.
	m.printBaseExpansionResult(ret)
	return nil
}

