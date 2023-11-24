package main

import (
	"fmt"
	"time"
)

func mySleep(time_ int) {
	dur := time_ * int(time.Second)
	<-time.After(time.Duration(dur))
}

func main() {
	fmt.Println("Start")
	mySleep(2)
	fmt.Println("End")
}
