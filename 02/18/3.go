package main

import (
	"fmt"
	"time"
)

type ReplicationManager struct {
	// ...
}

func (rm *ReplicationManager) watchControllers(controllerEvents chan string) {
	for {
		select {
		case event, open := <-controllerEvents:
			if !open {
				fmt.Println("Controller events channel closed.")
				return
			}
			fmt.Printf("Received controller event: %s\n", event)
		}
	}
}

func main() {
	rm := &ReplicationManager{}
	controllerEvents := make(chan string)

	go rm.watchControllers(controllerEvents)

	// 模拟发送副本控制器变更事件
	controllerEvents <- "create"
	controllerEvents <- "update"
	controllerEvents <- "delete"

	// 等待事件处理完成
	time.Sleep(2 * time.Second)

	// 关闭通道以结束watchControllers循环
	close(controllerEvents)

	// 等待watchControllers方法输出关闭信息
	time.Sleep(1 * time.Second)
}
