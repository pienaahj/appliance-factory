package main

import (
	"fmt"

	"github.com/pienaahj/appliance-factory/appliances"
)

func main() {
	// Uses the class factory to create an appliance of the requested type
	// Requests the user to enter the appliance type
	fmt.Println("Enter preffered appliance type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	// use fmt.scan to retrieve the user input
	var myType int
	fmt.Scan(&myType)

	myAppliance, err := appliances.CreateAppliance(myType)

	//if no errors, start the appliance and then print its putpose
	if err == nil {
		myAppliance.Start()
		fmt.Println((myAppliance.GetPurpose()))
	} else {
		// if error, print it
		fmt.Println(err)
	}
	
}