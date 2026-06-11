package main

import "fmt"

/*
	 an interface is defined by implementing a set of methods. If a struct does not have any methods to implement but needs to implement an interface,
	an empty struct can be used as a placeholder to indicate that the struct implements the interface.

	In above code, we define an empty interface myInterface, and a struct myStruct with a String() method that returns myStruct.
	Then, we implement myInterface by assigning myStruct{} to the interface variable i, making myStruct a type that implements myInterface.

	In this example, since myInterface is an empty interface with no methods, myStruct as an implementation of myInterface does not need to
	 implement any methods. However, since myStruct implements the String() method of the fmt.Stringer interface, we can print the value
	 of i using fmt.Println and get the string representation of myStruct


*/

type myInterface interface{}

type myStruct struct{}

func (ms myStruct) String() string {
	return "myStruct"
}

func main() {
	var i myInterface = myStruct{}
	fmt.Println(i)
}
