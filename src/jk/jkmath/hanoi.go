package jkmath

import (
	"fmt"
)

type Hanoi struct {
	n     int
	a     string
	b     string
	c     string
	step  []string
	count int64
}

func NewHanoi(n int) *Hanoi {
	return &Hanoi{
		n:     n,
		a:     "A",
		b:     "B",
		c:     "C",
		count: 0,
	}
}

func (h *Hanoi) Do() {
	h.do(h.n, h.a, h.b, h.c)
}

func (h *Hanoi) do(n int, a, b, c string) {
	if n == 1 {
		h.count += 1
		str := fmt.Sprintf("%s -> %s in step %d", a, c, h.count)
		h.step = append(h.step, str)
	} else {
		h.do(n-1, a, c, b)
		h.count += 1
		str := fmt.Sprintf("%s -> %s in step %d", a, c, h.count)
		h.step = append(h.step, str)
		h.do(n-1, b, a, c)
	}
}

func (h *Hanoi) Print() {
	for _, v := range h.step {
		fmt.Printf("do: %v\n", v)
	}
}
