package main

import (
	"fmt"
	"time"
)

type Worker struct {
	duration time.Duration
}

func NewWorker(duration time.Duration) *Worker {
	return &Worker{
		duration: duration,
	}
}

func (w *Worker) Run() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(w.duration)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func main() {
	worker := NewWorker(5 * time.Second)
	worker.Run()
}
