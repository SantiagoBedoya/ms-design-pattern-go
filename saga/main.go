package main

import (
	"fmt"
)

func action1(data string) error {
	fmt.Println("Executing action 1: ", data)
	return nil
}
func action2(data string) error {
	fmt.Println("Executing action 2: ", data)
	return nil
}
func action3(data string) error {
	fmt.Println("Executing action 3: ", data)
	return fmt.Errorf("error in action 3")
}
func compensateAction1(data string) error {
	fmt.Println("compensating action 1: ", data)
	return nil
}
func compensateAction2(data string) error {
	fmt.Println("compensating action 2: ", data)
	return nil
}
func compensateAction3(data string) error {
	fmt.Println("compensating action 3: ", data)
	return nil
}
func performSagaActions(data string) {
	fmt.Println("Starting saga")

	err := action1(data)
	if err != nil {
		compensateAction1(data)
		fmt.Println("failed saga")
		return
	}

	err = action2(data)
	if err != nil {
		compensateAction2(data)
		compensateAction1(data)
		fmt.Println("failed saga")
		return
	}

	err = action3(data)
	if err != nil {
		compensateAction3(data)
		compensateAction2(data)
		compensateAction1(data)
		fmt.Println("saga failed")
		return
	}

	fmt.Println("saga completed successfully")
}

func main() {
	data := "Test Data"
	performSagaActions(data)
}
