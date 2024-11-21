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
)

func main() {
	fmt.Println(package_1.Random())
}
