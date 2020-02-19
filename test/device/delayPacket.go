package device

//NewDelayPacket 创建一个延迟数据包
func NewDelayPacket(size int, src []byte) *DelayPacket {
	dp := &DelayPacket{
		_ptr:  make([]byte, size),
		_size: size,
	}

	if src != nil {
		copy(dp._ptr, src)
	}

	return dp
}

//DelayPacket 延迟数据包
type DelayPacket struct {
	_ptr  []byte
	_size int
	_ts   uint32
}
