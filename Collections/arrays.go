package Collections

import "fmt"

func ArrayTest() {

	var numbers [5]int

	fmt.Println(numbers)

	numbers[4] = 20

	fmt.Println(numbers)

	////

	fruits := [4]string{"Apple", "Banana", "Orange", "grapes"}
	fmt.Println(fruits)

	fmt.Println(fruits[2])

	for i := 0; i < len(fruits); i++ {

		fmt.Println("Elemet at index", i, ":", fruits[i])
	}

	// underscore (_) is blank identifier used to store unused values
	for index, val := range fruits {

		fmt.Printf("Index: %d, Value %v\n", index, val)

	}

	var matix [3][3]int = [3][3]int{{1, 2, 3}, {2, 4, 6}, {4, 6, 7}}

	fmt.Println(matix)

	//Pointer Array

	originalArray := [4]string{"Apple", "Banana", "Orange", "grapes"}

	var copiedArray *[4]string
	copiedArray = &originalArray

	copiedArray[1] = "Pinapple"
	fmt.Println(originalArray)

}
