package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Define the regular expression pattern.
	regexPattern := `^\d+\s+(\S+)`
	re := regexp.MustCompile(regexPattern)

	// Read from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if match != nil && len(match) > 1 {
			domain := match[1]
			fmt.Println(domain)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
	}
}
