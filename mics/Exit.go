package mics

import (
	"fmt"
	"os"
)

func ExitTes() {

	defer fmt.Println("this wont print because the application is already close by os.ex")
	fmt.Println("This will print")

	//this is used to exit application
	os.Exit(1)
}
