package jkmath

// package main

import (
	"errors"
	"fmt"
	"jk/jklog"
)

type Modular struct {
	n   int
	b   int
	ret []int
}

func New() *Modular {
	mod := Modular{}
	return &mod
}

func (m *Modular) Print() {
	jklog.L().Info("The result is : ")
	rlen := len(m.ret)
	for i := rlen - 1; i >= 0; i-- {
		fmt.Printf("%d", m.ret[i])
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
	m.ret = m.ret[:0]
	for {
		// Mod last, exit
		if q == 0 {
			break
		}
		quo := q % b
		m.ret = append(m.ret, quo)
		q = q / b
	}
	return nil
}
