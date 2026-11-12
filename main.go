package main

import "fmt"

func main() {
	var choice int

	fmt.Println("===== Hashing Labs W06 =====")
	fmt.Println("0. Lab 0: Proof the Hash Program")
	fmt.Println("1. Lab 1: MD5 Password Cracking")
	fmt.Println("2. Lab 2: SHA1 Password Cracking")
	fmt.Println("3. Lab 3: SHA512 Password Cracking")
	fmt.Print("Choose lab: ")
	fmt.Scanln(&choice)
	fmt.Println()

	switch choice {
	case 0:
		lab0()
	case 1:
		lab1()
	case 2:
		lab2()
	case 3:
		lab3()
	default:
		fmt.Println("Invalid choice")
	}
}
