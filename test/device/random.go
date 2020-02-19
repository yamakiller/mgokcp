package device

import (
	"math/rand"
)

//Random 均匀分布随机数
type Random struct {
	_size  int
	_seeds []int
}

//Next 获取一个随机数
func (slf *Random) Next() int {
	var x, i int
	if len(slf._seeds) == 0 {
		return 0
	}

	if slf._size == 0 {
		for i := 0; i < len(slf._seeds); i++ {
			slf._seeds[i] = i
		}
		slf._size = len(slf._seeds)
	}
	i = rand.Intn(slf._size)
	x = slf._seeds[i]
	slf._size--
	slf._seeds[i] = slf._seeds[slf._size]
	return x
}
