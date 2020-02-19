package device

import (
	"math/rand"

	"github.com/yamakiller/magicLibs/util"
)

//NewLatencySimulator 10,60,125, 1000
func NewLatencySimulator(lostrate, rttmin, rttmax, nmax int) *LatencySimulator {
	l := &LatencySimulator{
		_r12:      Random{_seeds: make([]int, 100)},
		_r21:      Random{_seeds: make([]int, 100)},
		_current:  uint32(util.Timestamp()),
		_lostrate: lostrate / 2,
		_rttmax:   rttmax / 2,
		_rttmin:   rttmin / 2,
		_nmax:     nmax,
	}

	return l
}

//LatencySimulator 网络延迟模拟器
type LatencySimulator struct {
	Tx1       int
	Tx2       int
	_current  uint32
	_lostrate int
	_rttmin   int
	_rttmax   int
	_nmax     int
	_p12      []*DelayPacket
	_p21      []*DelayPacket
	_r12      Random
	_r21      Random
}

//Clear 清理资源
func (slf *LatencySimulator) Clear() {
	slf._p12 = make([]*DelayPacket, 0)
	slf._p21 = make([]*DelayPacket, 0)
}

//Send 发送数据
func (slf *LatencySimulator) Send(peer int, data []byte, size int) {
	if peer == 0 {
		slf.Tx1++
		if slf._r12.Next() < slf._lostrate {
			return
		}

		if len(slf._p12) >= slf._nmax {
			return
		}
	} else {
		slf.Tx2++
		if slf._r21.Next() < slf._lostrate {
			return
		}
	}

	if len(slf._p21) >= slf._nmax {
		return
	}

	pkt := NewDelayPacket(size, data)
	slf._current = uint32(util.Timestamp())
	delay := uint32(slf._rttmin)
	if slf._rttmax > slf._rttmin {
		delay += uint32(rand.Intn(slf._rttmax - slf._rttmin))
	}
	pkt._ts = slf._current + delay
	if peer == 0 {
		slf._p12 = append(slf._p12, pkt)
	} else {
		slf._p21 = append(slf._p21, pkt)
	}
}

//Recv 接收数据
func (slf *LatencySimulator) Recv(peer int, data []byte, maxSize int) int {
	var it *DelayPacket
	if peer == 0 {
		if len(slf._p21) == 0 {
			return -1
		}
		it = slf._p21[0]
	} else {
		if len(slf._p12) == 0 {
			return -1
		}
		it = slf._p12[0]
	}
	pkt := it
	slf._current = uint32(util.Timestamp())
	if slf._current < pkt._ts {
		return -2
	}
	if maxSize < pkt._size {
		return -3
	}
	if peer == 0 {
		slf._p21 = slf._p21[1:]
	} else {
		slf._p12 = slf._p12[1:]
	}
	maxSize = pkt._size
	copy(data, pkt._ptr)
	return maxSize
}
