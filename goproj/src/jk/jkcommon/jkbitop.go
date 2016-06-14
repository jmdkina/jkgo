package jkcommon

import (
	// "jk/jklog"
)

func JKBitValue(value, start, cout uint) uint {
	ret := uint(0)
	var i uint
	i = 0
	var j uint
	j = 0 
    for i = start - cout + 1; i <= start; i++ {
        v := uint(value & (1 << i))
        if v != 0 {
        	v = uint(1)
        }
        ret += v * 2^j
        j++
    }
    return ret
}