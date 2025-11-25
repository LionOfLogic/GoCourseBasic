package mics

import "fmt"

func RecoverMain() {

	RecoverTest()
	fmt.Println("Returned from process")

}

func RecoverTest() {

	//recover() this fucntion is built in to recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverdd", r)
		}

	}()

	fmt.Println("Start Process")
	panic("Something happended!")

	fmt.Println("End Process")

}
