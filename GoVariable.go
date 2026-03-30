package main

import (
	"fmt"
)

var gVar = "Golobal"

func abcc() {
	var lVar = "Lolobal"
	var a int = 1
	fmt.Println(a)
	a = 4
	fmt.Println(gVar)
	fmt.Println(lVar)
}

func Amain() {

	var a int = 1
	fmt.Println(a)
	a = 4
	fmt.Println(gVar)
}
