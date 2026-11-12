package main

import (
	"fmt"

	"Hashing/utils/crack"
)

func lab2() {
	target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	wordlist := "nord_vpn.txt"
	verbose := "verbose_lab2.txt"

	fmt.Println("Lab2 SHA1 Cracking ...")
	_, _ = crack.CrackSHA1(target, wordlist, verbose)
}
