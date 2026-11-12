package main

import (
	"fmt"

	"Hashing/utils/crack"
)

func lab1() {
	target := "6a85dfd77d9cb35770c9dc6728d73d3f"
	wordlist := "nord_vpn.txt"
	verbose := "verbose_lab1.txt"

	fmt.Println("Lab1 MD5 Cracking ...")
	_, _ = crack.CrackMD5(target, wordlist, verbose)
}
