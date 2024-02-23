package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	utils "github.com/IamLucif3r/Go-Hash-Matcher/utils"
)

var wg sync.WaitGroup

func checkPassword(value, search, hashType, salt string, results chan bool) {
	defer wg.Done()
	hashedPassword := utils.CryptBytes(hashType, salt, []byte(value))
	if hashedPassword == search {
		fmt.Printf("[✅] Hash Matched! Password Found : %s, hash: %s\n", value, hashedPassword)
		results <- true
	}
}

func main() {
	hashType := "SHA1"
	salt := "d"
	// Enter Hash you want to decrypt
	search := "$SHA1$d$uP0_QaVBpDWFeo8-dRzDqRwXQ2I"
	// Provide path for a list of plain text passwords
	wordlist := "/usr/share/wordlists/rockyou.txt"

	go utils.ShowLoadingSpinner()

	passwordListFile, err := os.Open(wordlist)
	if err != nil {
		fmt.Printf("[❗️] Error opening wordlist: %v\n", err)
		return
	}
	defer passwordListFile.Close()

	scanner := bufio.NewScanner(passwordListFile)
	results := make(chan bool)

	for scanner.Scan() {
		value := scanner.Text()
		wg.Add(1)
		go checkPassword(value, search, hashType, salt, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	select {
	case <-results:
	case <-time.After(10 * time.Second):
		fmt.Println("[❌] Password not found in the given timeout.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("[❗️] Error reading wordlist: %v\n", err)
	}
}
