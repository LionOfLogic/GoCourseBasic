package mics

import "fmt"

func PanicTest() {

}

func process(input int) {

	defer fmt.Println("This is will registe")

	if input < 0 {
		panic("input must be non -zero")
		defer fmt.Println("This wont because panic exit the fucntion already will registe")

	}

}
