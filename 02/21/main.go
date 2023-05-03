package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个异步任务，使其在3秒后返回字符串 "Hello, Async!"
	asyncTask := makeAsync(func() interface{} {
		time.Sleep(3 * time.Second)
		return "Hello, Async!"
	})

	// 使用waitForResult函数等待异步任务完成，并输出结果
	result, err := waitForResult(asyncTask, 5*time.Second)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func makeAsync(fn func() interface{}) <-chan interface{} {
	channel := make(chan interface{}, 1)
	go func() {
		channel <- fn()
	}()
	return channel
}

func waitForResult(out <-chan interface{}, timeout time.Duration) (interface{}, error) {
	select {
	case result := <-out:
		return result, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("Timed out waiting for result.")
	}
}
