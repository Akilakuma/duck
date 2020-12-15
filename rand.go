package duck

import (
	"bytes"
	"encoding/binary"
	"os/exec"
)

var (
	// DefaultRandInstance 預設的亂數實體
	DefaultRandInstance *RandManager

	// MaxRandQ 亂數Channel最大的數量
	MaxRandQ int = 3000000
)

// RandManager 亂數產生業務
type RandManager struct {
	randQueue  chan int
	production int // 總產生量

}

// New 新的亂數產生器
func New(maxStorageRandNum int) *RandManager {
	if MaxRandQ != 0 {
		MaxRandQ = maxStorageRandNum
	}
	r := &RandManager{
		randQueue: make(chan int, MaxRandQ),
	}
	DefaultRandInstance = r
	for i := 0; i < 10; i++ {
		go DefaultRandInstance.KeepGenerateRandNum()
	}
	return r
}

// KeepGenerateRandNum 持續產生亂數
func (r *RandManager) KeepGenerateRandNum() {
	for {
		numSlice := r.RandGenerate()
		for _, num := range numSlice {
			r.randQueue <- num
			r.production++
		}
	}
}

// RandGenerate 亂數產生
func (r *RandManager) RandGenerate() []int {
	var (
		out       bytes.Buffer
		isSuccess bool
	)

	for !isSuccess {
		//本機mac開發無/dev/hwrng
		cmd := exec.Command("head", "-c", "600000", "/dev/urandom")
		cmd.Stdout = &out
		err2 := cmd.Run()
		if err2 != nil {
			// fmt.Println("ErrSystemGenerateRand" + err2.Error())
			continue
		}
		isSuccess = true
	}
	n := bytes.Split(out.Bytes(), []byte(" "))

	var numSlice []int
	for _, v := range n {
		if len(v) > 4 {
			randNum := binary.BigEndian.Uint32(v)
			numSlice = append(numSlice, int(randNum))
		}
	}

	return numSlice
}

// GetRand 取得亂數
func (r *RandManager) getRand() int {
	n, ok := <-r.randQueue
	if ok {
		return n
	}
	return 0
}

// GetRandStorageNum 取得目前現有亂數庫存數量
func (r *RandManager) GetRandStorageNum() int {

	return len(r.randQueue)
}

// GetRandUnderRange 取得mod 條件的亂數
func (r *RandManager) GetRandUnderRange(modNum int) int {

	n := r.getRand()
	if modNum == 0 {
		return 0
	}

	randNum := n % modNum
	return randNum
}

// GetRandBetweenRange 取得指定 min ~ max的亂數
func (r *RandManager) GetRandBetweenRange(min, max int) int {

	num := r.GetRandUnderRange(max + 1)

	// 數量最少要達到低標
	if num < min {
		num = min
	}
	return num
}
