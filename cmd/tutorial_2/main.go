package main

import (
	"fmt"
	"unicode/utf8"
)

// variables
func main() {
	var intNum int
	intNum = 14
	fmt.Println("Num: ", intNum)

	var floatNum float64 = 45.01
	fmt.Println("Float: ", floatNum)

	//convert types
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	result := floatNum32 + float32(intNum32)
	fmt.Println("result: ", result)

	intNum1 := 3
	intNum2 := 2

	fmt.Println("Answer using /: ", intNum1/intNum2)
	fmt.Println("Answer using %: ", intNum1%intNum2)

	var myString string = "Hello There"
	fmt.Println(myString)

	//counting the number of character in string not using len but a another package
	//this gamma symbols count as 2 character
	fmt.Println(len("γ"))

	//use this instead
	fmt.Println(utf8.RuneCountInString("γ"))

	var myRune rune = 'a'
	fmt.Println("Rune?: ", myRune)

	var myBool bool = true
	fmt.Println("Bool: ", myBool)

	const pi float32 = 3.1415
}
