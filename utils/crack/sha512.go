package crack

import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"os"
)

// CrackSHA512 reads the wordlist line by line, hashes each word with SHA512,
// and compares it to target. It writes a verbose log to verbosePath and
// returns the cracked password (empty if not found).
func CrackSHA512(target, wordlistPath, verbosePath string) (string, error) {
	// Open the wordlist for reading.
	file, err := os.Open(wordlistPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create the verbose output file.
	out, err := os.Create(verbosePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	fmt.Fprintf(out, "Target SHA512 : %s\n", target)
	fmt.Fprintf(out, "Wordlist      : %s\n\n", wordlistPath)

	found := ""
	attempts := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		attempts++

		hash := fmt.Sprintf("%x", sha512.Sum512([]byte(word)))

		// Log every attempt to the verbose file.
		fmt.Fprintf(out, "[%d] %s => %s\n", attempts, word, hash)

		if hash == target {
			found = word
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	fmt.Fprintf(out, "\nAttempts: %d\n", attempts)
	if found != "" {
		fmt.Fprintf(out, "PASSWORD FOUND: %s\n", found)
		fmt.Printf("PASSWORD FOUND: %s (after %d attempts)\n", found, attempts)
	} else {
		fmt.Fprintln(out, "PASSWORD NOT FOUND")
		fmt.Printf("PASSWORD NOT FOUND (after %d attempts)\n", attempts)
	}

	return found, nil
}
