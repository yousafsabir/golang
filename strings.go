package main

import "fmt"

func StringsDeepDive() {
	myString := "résumé"

	// 1- We can index a string just like a regular Array
	fmt.Println(myString[0])
	// But it doesn't print out the character, but rather an integer
	fmt.Printf("%v of type %T", myString[0], myString[0])

}
