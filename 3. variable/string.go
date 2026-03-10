package main

import (
	"fmt"
	"strings"
)

func main() {
	nama := "Ilham"
	message := "Selamat Datang di Aplikasi Go"

	paragraf := `Halo, ini adalah contoh multi
	 line string di Go`
	fmt.Println("Nama : ", nama)
	fmt.Println("Message : ", message)
	fmt.Println("Paragraf : \n", paragraf)


	text:= "Halo dunia"

	// ubah jadi huruf kecil
	fmt.Println("Lowercase : ", strings.ToLower(text))
	// ubah jadi huruf besar
	fmt.Println("UpperCase : ", strings.ToUpper(text))
	// cek apakah string dimulai dari kata tertentu 
	fmt.Println("StarsWith Halo ? : ", strings.HasPrefix(text, "Halo"))
	// cek apakah string mengandung kata tertentu
	fmt.Println("Contains Dunia ? : ", strings.Contains(text, "dunia"))
	//mengganti bagian string
	newText := strings.ReplaceAll(text, "dunia", "Golang")
	fmt.Println("Replace, ",newText)
}