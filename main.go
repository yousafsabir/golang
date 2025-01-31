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
	"fmt"
	"go_tutorial/package_1"
	"unicode/utf8"
	"utf8"
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
}
