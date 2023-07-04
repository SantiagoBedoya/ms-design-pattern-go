package main

import (
	"fmt"
	"sync"
	"time"
)

type Bulkhead struct {
	semaphore chan struct{}
}

func NewBulkhead(concurrency int) *Bulkhead {
	return &Bulkhead{
		semaphore: make(chan struct{}, concurrency),
	}
}

func (b *Bulkhead) Execute(task func()) {
	b.semaphore <- struct{}{}
	go func() {
		defer func() {
			<-b.semaphore
		}()
		task()
	}()
}

func main() {
	bulkhead := NewBulkhead(2)

	var wg sync.WaitGroup

	tasks := []func(){
		func() {
			fmt.Println("Starting Task 1")
			time.Sleep(1 * time.Second)
			fmt.Println("Task 1 terminated")
			wg.Done()
		},
		func() {
			fmt.Println("Starting Task 2")
			time.Sleep(1 * time.Second)
			fmt.Println("Task 2 terminated")
			wg.Done()
		},
		func() {
			fmt.Println("Starting Task 3")
			time.Sleep(1 * time.Second)
			fmt.Println("Task 3 terminated")
			wg.Done()
		},
	}

	for _, task := range tasks {
		wg.Add(1)
		bulkhead.Execute(task)
	}

	wg.Wait()
}
