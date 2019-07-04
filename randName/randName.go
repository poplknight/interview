package randName

import (
	"fmt"
	"strconv"
	"sync"
)
type randName struct {
	i uint64
	bit uint8
	mu sync.Mutex
}

var rn *randName = &randName{0,4,sync.Mutex{}}

func GetRandName(n int) string {
	var i uint64
	func(){
		rn.mu.Lock()
		defer rn.mu.Unlock()
		i = rn.i
		rn.i ++
	}()

	return	fmt.Sprintf("%0"+strconv.Itoa(n)+"s",strconv.FormatUint(rn.pseudo_encrypt(i),10))
}

func setRandNameBit(n int) {
	//这里要把10进制的位数，换算成2进制，二进制位数必须为偶数
	return
}

func (rn *randName)pseudo_encrypt(value uint64)uint64 {
	var b uint64 = 1 << rn.bit - 1
	var l, r uint64 = (value >> rn.bit) & b, value & b

	for i := 0; i < 3; i ++{
		l, r = r, l ^ uint64(uint64((float64((1366 * r + 150889) % 714025)/ 714025.0) * 32767 * 32767) & b)
	}
	return ((l << rn.bit) + r)
}