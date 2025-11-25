package Collections

import "fmt"

func MapTest() {

	//VAR mapVriable map[keyType]ValueType
	fmt.Println("Maps Test")

	myMap := make(map[string]int)

	myMap["key1"] = 1

	fmt.Println(myMap)

	//a Map return value and if key exits bool
	value, isKeyExits := myMap["key1"]

	fmt.Println(value)
	fmt.Println(isKeyExits)
}
