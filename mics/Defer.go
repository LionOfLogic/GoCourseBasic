package mics

import "fmt"

//defer is use to deley the call mainly use for clean up
func DeferTest() {

	val := 5

	// this will registerd but wont print it will print at last with value 5
	defer fmt.Println("This will be call at last ", val)

	val += 1
	for i := 0; i < 3; i++ {

		println("Check :", i)
	}
	fmt.Println(val)
}
