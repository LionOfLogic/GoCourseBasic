package main

import (
	"fmt"
	col "gocoursebasic/collectionm"
)

func main() {

	fmt.Println("Hello world!")
	//ifelseTest()
	//forloopTest()
	//switchTest()
	//whileTest()
	//col.ArrayTest()
	//col.SlicesTest()
	col.MapTest()
}

func ifelseTest() {

	var num int
	fmt.Println("Give a number input")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("Its a positive number")
	} else if num < 0 {
		fmt.Println("It a negavtive number")
	} else {
		fmt.Println("Its a Zero")
	}

}

func forloopTest() {

	var num int
	fmt.Println("Give a input")

	fmt.Scanln(&num)

	for i := 0; i <= num; i++ {
		fmt.Println(i)
	}

}

func switchTest() {

	var num int
	fmt.Println(`press 1 for English
	press 2 for Hindi
	press 3 for Marathi`)

	fmt.Scanln(&num)

	switch num {
	case 1:
		println("You chose English")
	case 2:
		println("You chose Hindi")
	case 3:
		println("You chose Marathi")
	default:
		println("Invalid choice")
	}

}

func whileTest() {

	var num int
	fmt.Println("Enter the input")

	fmt.Scanln(&num)

	for num > 0 {

		fmt.Printf("The number is %d\n", num)
		num -= 1

	}

}
