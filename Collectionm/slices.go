package collectionm

import (
	"fmt"
	"slices"
)

func SlicesTest() {

	var numbers1 = []int{1, 2, 3}

	number2 := []int{9, 8, 7}

	slice := make([]int, 5)

	fmt.Println(number2, numbers1, number2, slice)

	//slice extration

	number3 := number2[:2]

	fmt.Println(number3)

	number4 := number2[1:2]

	fmt.Println(number4)

	//append
	number4 = append(number4, 8, 9, 10)
	fmt.Println(number4)

	//copy slice
	slicCopy := make([]int, len(number3))
	copy(slicCopy, number3)
	fmt.Println(slicCopy)

	//iterate
	for _, v := range number4 {

		fmt.Println(v)
	}

	//compare slices
	if slices.Equal(number3, slicCopy) {
		fmt.Println("slices is equal to sliceCopy")
	}

	//multidimensional Slices

	towDSlice := make([][]int, 3)

	for i := 0; i < 3; i++ {

		interLen := i + 1

		towDSlice[i] = make([]int, interLen)
		for j := 0; j < interLen; j++ {
			towDSlice[i][j] = i + j
		}

	}

	fmt.Println(towDSlice)

	fmt.Println("the capacity of slice is ", cap(towDSlice))

}
