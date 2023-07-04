package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	retryCount := 3

	for i := 0; i < retryCount; i++ {
		err := doSomething()
		if err == nil {
			break
		}

		fmt.Printf("Error: %s\n", err)
		waitTime := time.Second * 2
		fmt.Printf("Waiting %s before the next try\n", waitTime)
		time.Sleep(waitTime)
	}
}

func doSomething() error {
	n := rand.Intn(10)
	if n != 5 {
		return fmt.Errorf("the number is not 5")
	}
	return nil
}
