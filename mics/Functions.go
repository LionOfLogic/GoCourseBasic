package mics

import "fmt"

func FunctionMain() {

	fmt.Println("This is function")

	greet := func() {

		fmt.Println("This is anonymous fucntion")
	}

	greet()

}

//first class citizen function

func applyOperation(x int, y int, operation func(int, int) int) int {

	return operation(x, y)
}
