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

	fmt.Println("Nilai Ujian: ", text, " sudah menjadi String")

	//String to int
	var text2 string = "100a"
	number, err := strconv.Atoi(text2)
	if err != nil {
		fmt.Println("Error message: ", err.Error())
	} else {
		fmt.Println("\nAngka: ", number, " sudah menjadi Int")
	}

	//Boolean to String
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("\nBoolean ke String menjadi : ",str)
	//String to Boolean
	val, _ := strconv.ParseBool("true")
	fmt.Println("String ke boolean: ", val)
}