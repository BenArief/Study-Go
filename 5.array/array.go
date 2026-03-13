package main

import "fmt"

func main1() {
	fmt.Println("Array in Golang")
	var angka [3]int = [3]int{10, 20, 30}
	fmt.Println(angka)
	fmt.Println(angka[1])
	angka[1] = 80
	fmt.Println("Ini adalah value index 1: ", angka[1])
}