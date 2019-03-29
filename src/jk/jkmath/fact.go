package jkmath

func Fact(i int) int64 {
	if i == 0 {
		return int64(1)
	} else {
		return int64(int64(i) * Fact(i-1))
	}
}
