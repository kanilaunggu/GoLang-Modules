package main

import (
	"fmt"
)

func main() {
	//Golang memungkinkan Anda untuk menetapkan nilai
	//ke beberapa variabel dalam satu baris.
	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}
