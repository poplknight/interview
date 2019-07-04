package randName

import (
	"strconv"
	"sync"
)
type randName struct {
	i uint64
	mu sync.Mutex
}

const uint32Max  = 0xffffffff

var rn *randName = &randName{0,sync.Mutex{}}

func GetRandName() string {
	var i uint64
	func() {
		rn.mu.Lock()
		defer rn.mu.Unlock()
		i = rn.i
		rn.i ++
	}()

	return strconv.FormatUint(rn.pseudo_encrypt(i),10)
}

func (rn *randName)pseudo_encrypt(value uint64)uint64 {
	var l, r uint64 = (value >> 32) & uint32Max, value & uint32Max

	for i := 0; i < 3; i ++{
		l, r = r, l ^ uint64((float64((1366 * r + 150889) % 714025)/ 714025.0) * 32767 * 32767)
	}
	return ((l << 32) + r)
}