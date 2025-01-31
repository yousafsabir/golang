package main

// 1- Go Modules & Packages:
// 		Packages:
// 				- a folder containing one or more .go source files for the purpose of splitting code.
// 				  one file can access variables & functions from another file of the same packages.
// 				- all files in a package must be of the same package.
//  			- a package is part of a module.
//
//  			here I have package_1 of the go_tutorial module
//
// 		Modules:
// 				- a module is a like a complete go project with a go.mod file
//
// 				Just like this module is named go_tutorial

import (
	"errors"
	"fmt"
	"go_tutorial/package_1"
	"unicode/utf8"
)

func main() {
	fmt.Println(package_1.Random())

	// Variables & Data Types
	// Integers
	var intNum int = 345
	fmt.Println(intNum)
	// we can also specify the exact memory of each int variable by using
	// int8, int16, int32, int64
	// We also have unsigned integers that only store positive values
	// uint, uint8, uint16, uint32, uint64

	// Floats
	// we have two data types float32 & float64 for storing float numbers
	// there is no float type
	var floatNum float64 = 123456789.9
	fmt.Println(floatNum)

	// Arithematic Operations
	// can only be performed with same operand data types
	// the result would be the in the same type as of operands
	var float1 float64 = floatNum + float64(intNum)
	fmt.Println(float1)

	// Strings
	// string data type is used to declare strings
	var myString string = "Hello"
	fmt.Println(myString)
	// string can use double quotes " for single line or back ticks ` for multi line
	myString = `Hello
World`
	// we can also concatenate strings
	myString = "Hello" + " " + "\nWorld"

	// String Length
	// for string length, we can use the built in len() function just like python
	fmt.Println(len(myString))
	// But there is a catch, len() is fine for regular characters. But for special
	// characters, its not accurate. That's because it counts the underlying bytes
	// each character has, not the characters themselves.
	// for that we can use RuneCountInString func from utf8 package
	fmt.Println(utf8.RuneCountInString(myString))

	// Rune
	// Now rune is a single character data type just like char in c++
	var myRune rune = 'a'
	fmt.Println(myRune)
	// They're rarely used

	// Booleans
	// we have a bool data type with two values true & false as usual
	var myBoolean bool = true
	fmt.Println(myBoolean)
	myBoolean = false

	// Default Values
	// If variable not initialised, Go sets
	// 0 for all ints, uints, floats & runes
	// '' for strings
	// false for booleans

	// Short Hand Declaration
	// we can omit the type if we initilise a variable on its creation
	// Go will automatically infer its type
	var myString2 = "HelloWorld"
	fmt.Println(myString2)
	// Or we can drop the var keyword & type together with := operator
	myString3 := "Hello World"
	fmt.Println(myString3)
	// we can also declare multiple vars in one line
	var num1, num2 = 1, 2
	num3, num4 := 3, 4

	fmt.Println(num1, num2, num3, num4)

	// Constants
	// constants are declared the same way as variables except
	// Their value can't be changed later (obviosly)
	// They can't be declared uninitialised
	const myConst = "a constant"
	fmt.Println(myConst)

	printMe("Hello")

	var result, reminder, error = intDivisionWithReminder(5, 3)

	// Control Structures
	// If Else
	if error != nil {
		fmt.Printf(error.Error())
	} else if reminder == 0 {
		fmt.Printf("Result is %v", result)
	} else {
		fmt.Printf("Result is %v & the reminder is %v", result, reminder)
	}
	// Switch Case
	switch {
	case error != nil:
		fmt.Printf(error.Error())
	case reminder == 0:
		fmt.Printf("Result is %v", result)
	default:
		fmt.Printf("Result is %v & the reminder is %v", result, reminder)

	}
	// Conditional Switch Case
	// Applies on a specific value
	switch reminder {
	case 0:
		fmt.Println("The Division was exact")
	case 1, 2:
		fmt.Println("The Division was close")
	default:
		fmt.Println("The Division was not close")
	}
}

// Functions
// are declared with func keyword
func printMe(printVal string) {
	fmt.Println("Function Printed")
	fmt.Println(printVal)
}

func intDivision(numerator int, denominator int) int {
	return numerator / denominator
}

// can also return multiple value at once
func intDivisionWithReminder(numerator int, denominator int) (int, int, error) {
	// if denominator is 0, then we get an error
	// we have to check if denominator is 0
	// and return the error if it's 0
	var err error // default value: nil
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		// Notice that we do have to return some value for
		// result & the reminder
		return 0, 0, err
	}
	// err is nil here
	return numerator / denominator, numerator % denominator, err
}
