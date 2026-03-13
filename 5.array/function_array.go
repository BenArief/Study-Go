package main

import "fmt"

func main() {
	number := [5]int{10, 20, 30, 40,50}

	fmt.Println("Jumlah elemen : ", len(number))
	fmt.Println(number[1])
	fmt.Println()

	for i, value := range number {
		fmt.Println("isi index ke-",i+1, " = ", value)
	}

	fmt.Println()

	arr1 := [3]int{1,2,3}
	arr2 := [3]int{4,5,6}
	fmt.Println("arr1 == arr2 : ", arr1==arr2)
	fmt.Println("arr1 != arr2 : ", arr1!=arr2)
}