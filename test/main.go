package main

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/mgokcp/mkcp"
	"github.com/yamakiller/mgokcp/test/device"
)

var (
	vnet *device.LatencySimulator
)

// 模拟网络：模拟发送一个 udp包
func udpOutput(buffer []byte, user interface{}) int32 {
	id := user.(int)
	vnet.Send(id, buffer, len(buffer))
	return 0
}

func getTime() uint32 {
	return uint32(util.Timestamp() & 0xffffffff)
}

//测试用例
func test(mode int) {
	// 创建模拟网络：丢包率10%，Rtt 60ms~125ms
	vnet = device.NewLatencySimulator(10, 60, 125, 1000)
	// 创建两个端点的 kcp对象，第一个参数 conv是会话编号，同一个会话需要相同
	// 最后一个是 user参数，用来传递标识
	kcp1 := mkcp.New(0x11223344, 0)
	kcp2 := mkcp.New(0x11223344, 1)

	kcp1.WithOutput(udpOutput)
	kcp2.WithOutput(udpOutput)

	current := getTime()
	slap := current + uint32(20)

	index := uint32(0)
	next := uint32(0)
	sumrtt := int64(0)
	count := 0
	maxrtt := 0

	// 配置窗口大小：平均延迟200ms，每20ms发送一个包，
	// 而考虑到丢包重发，设置最大收发窗口为128
	kcp1.WndSize(128, 128)
	kcp2.WndSize(128, 128)
	if mode == 0 {
		kcp1.NoDelay(0, 10, 0, 0)
		kcp2.NoDelay(0, 10, 0, 0)
	} else if mode == 1 {
		kcp1.NoDelay(0, 10, 0, 1)
		kcp2.NoDelay(0, 10, 0, 1)
	} else {
		kcp1.NoDelay(1, 10, 2, 1)
		kcp2.NoDelay(1, 10, 2, 1)

		kcp1.SetRxMinRto(10)
		kcp1.SetFastResend(1)
	}

	buffer := make([]byte, 2000)
	var hr int32
	ts1 := getTime()

	for {
		//fmt.Println("＃＃＃＃＃＃＃＃＃＃＃＃＃＃＃＃＃＃＃")
		time.Sleep(time.Millisecond)
		current = getTime()
		kcp1.Update(current)
		kcp2.Update(current)

		for ; current >= slap; slap += 20 {
			binary.LittleEndian.PutUint32(buffer, index)
			binary.LittleEndian.PutUint32(buffer[4:], current)
			index++
			// 发送上层协议包
			kcp1.Send(buffer, 8)
		}

		// 处理虚拟网络：检测是否有udp包从p1->p2
		for {
			hr = int32(vnet.Recv(1, buffer, 2000))
			if hr < 0 {
				break
			}

			kcp2.Input(buffer, hr)
		}
		// 处理虚拟网络：检测是否有udp包从p2->p1
		for {

			hr = int32(vnet.Recv(0, buffer, 2000))
			if hr < 0 {
				break
			}
			// 如果 p1收到udp，则作为下层协议输入到kcp1
			kcp1.Input(buffer, hr)
		}

		// kcp2接收到任何包都返回回去
		for {
			hr = kcp2.Recv(buffer, 10)
			// 没有收到包就退出
			if hr < 0 {
				break
			}
			// 如果收到包就回射
			kcp2.Send(buffer, hr)
		}

		// kcp1收到kcp2的回射数据
		for {
			hr = kcp1.Recv(buffer, 10)
			if hr < 0 {
				// 没有收到包就退出
				break
			}
			sn := binary.LittleEndian.Uint32(buffer)
			ts := binary.LittleEndian.Uint32(buffer[4:])
			rtt := current - ts
			if sn != next {
				fmt.Printf("ERROR sn %d<->%d\n", count, next)
				return
			}
			next++
			sumrtt += int64(rtt)
			count++

			if rtt > uint32(maxrtt) {
				maxrtt = int(rtt)
			}

			fmt.Printf("[RECV] mode=%d sn=%d rtt=%d\n", mode, sn, rtt)
		}

		if next > 1000 {
			break
		}
	}

	ts1 = getTime() - ts1

	mkcp.Free(kcp1)
	mkcp.Free(kcp2)

	names := make([]string, 3)
	names[0] = "default"
	names[1] = "normal"
	names[2] = "fast"

	fmt.Printf("%s mode result (%dms):\n", names[mode], int(ts1))
	fmt.Printf("avgrtt=%d maxrtt=%d tx=%d\n", sumrtt/int64(count), maxrtt, vnet.Tx1)
	fmt.Printf("press enter to next ...\n")
	var ch string
	fmt.Scanln(&ch)
	if ch != "" {

	}
}

func main() {
	fmt.Println("开始测试")

	test(0)
	test(1)
	test(2)
}
