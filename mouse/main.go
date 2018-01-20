package main

import (
	"github.com/go-vgo/robotgo"
	"time"
	"os"
	"fmt"
	"strconv"
)

var stop = true
var times = 10

func main() {
	fmt.Println("默认每秒10次，如果想要自定义请输入每秒次数")
	fmt.Println("点击c键开始，再次点击停止，如此循环")
	fmt.Println("要停止同时按下 ctrl 和 c")
	if len(os.Args) > 1 {
		var err error
		times, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("输入数字")
			os.Exit(1)
		}
	}

	c := make(chan bool)

	go func() {
		for {
			if robotgo.AddEvent("c") == 0 {
				if stop {
					stop = false
					fmt.Println("已启动", "每秒 ", times, " 次")
					go start()
				} else {
					fmt.Println("已停止")
					stop = true
				}
			}
		}

	}()

	<-c
}

func start() {
	for ; !stop; {
		robotgo.MouseClick("left")
		time.Sleep(time.Millisecond * (time.Duration(1000 / times)))
	}
}
