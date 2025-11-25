package mics

func VariMain() {

	//func functioName(param1 ...type)

	sum("For adition", 2, 3, 4, 5, 6)

}

// variatic paramter should be at last
func sum(returnString string, nums ...int) int {

	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}
