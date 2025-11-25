package mics

import "fmt"

func RangeTest() {

	message := "Hello world"

	//Range iterate over copy of given variable it doent change the value of original
	for i, v := range message {

		fmt.Printf("Index: %d, Rune %c\n", i, v)

	}

}
