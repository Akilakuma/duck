package main

import (
	"fmt"
	"time"

	"github.com/akilakuma/duck"
)

var random *duck.RandManager

func init() {
	random = duck.New(1000000)
}

func main() {

	// you need some time to wait rand generate from system
	time.Sleep(1 * time.Second)

	// print rand num in storage now
	fmt.Println(random.GetRandStorageNum())

	// get a rand num in a range
	fmt.Println(random.GetRandBetweenRange(1, 100))

	// get rand num under a num
	fmt.Println(random.GetRandUnderRange(1000))
}
