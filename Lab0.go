package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"fmt"
)

func proofMe(txt1, txt2 string) {

	names := []string{"MD5", "SHA1", "SHA256", "SHA512", "SHA3-256"}
	hashes := map[string][2]string{}

	// MD5
	md5A := fmt.Sprintf("%x", md5.Sum([]byte(txt1)))
	md5B := fmt.Sprintf("%x", md5.Sum([]byte(txt2)))
	hashes["MD5"] = [2]string{md5A, md5B}

	// SHA1
	sha1A := fmt.Sprintf("%x", sha1.Sum([]byte(txt1)))
	sha1B := fmt.Sprintf("%x", sha1.Sum([]byte(txt2)))
	hashes["SHA1"] = [2]string{sha1A, sha1B}

	// SHA256
	sha256A := fmt.Sprintf("%x", sha256.Sum256([]byte(txt1)))
	sha256B := fmt.Sprintf("%x", sha256.Sum256([]byte(txt2)))
	hashes["SHA256"] = [2]string{sha256A, sha256B}

	// SHA512
	sha512A := fmt.Sprintf("%x", sha512.Sum512([]byte(txt1)))
	sha512B := fmt.Sprintf("%x", sha512.Sum512([]byte(txt2)))
	hashes["SHA512"] = [2]string{sha512A, sha512B}

	// SHA3-256
	sha3A := fmt.Sprintf("%x", sha3.Sum256([]byte(txt1)))
	sha3B := fmt.Sprintf("%x", sha3.Sum256([]byte(txt2)))
	hashes["SHA3-256"] = [2]string{sha3A, sha3B}

	for _, algo := range names {
		vals := hashes[algo]
		match := "No Match!"
		if vals[0] == vals[1] {
			match = "Match!"
		}
		fmt.Printf("Hash (%s):\n  Output A = %s\n  Output B = %s\n  => %s\n\n", algo, vals[0], vals[1], match)
	}
}

func lab0() {
	fmt.Println("========Name + Hashing Program========")
	var txt1, txt2 string
	fmt.Print("Enter First String: ")
	fmt.Scanln(&txt1)
	fmt.Print("Enter Second String: ")
	fmt.Scanln(&txt2)

	proofMe(txt1, txt2)
}
