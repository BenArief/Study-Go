package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Int to Float
	var a int = 10
	var b float64 = float64(a)

	fmt.Println("Nilai a: ", a)
	fmt.Println("Nilai b: ", b)

	//Int to String
	var score int = 90
	var text string = strconv.Itoa(score)

	fmt.Println("Nilai Ujian: ", text)
}