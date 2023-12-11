package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var url string

	// Check if a URL argument is provided
	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		// No URL argument, prompt user for URL
		url = promptForURL()
	}

	// Process the URL
	processURL(url)
}

func promptForURL() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter download link: ")
	url, _ := reader.ReadString('\n')
	return strings.TrimSpace(url)
}

func processURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response:", err)
		return
	}
	bodyString := string(body)

	slug := extractBetween(bodyString, "\\\"slug\\\":\\\"", "\\\"")
	fileName := extractBetween(bodyString, "\\\"fileName\\\":\\\"", "\\\"")
	ext := fileName[strings.LastIndex(fileName, ".")+1:]

	fmt.Printf("https://qiwi.lol/%s.%s\n", slug, ext)
}

// extractBetween finds a substring between two delimiters
func extractBetween(s, start, end string) string {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return ""
	}
	startIndex += len(start)

	endIndex := strings.Index(s[startIndex:], end)
	if endIndex == -1 {
		return ""
	}

	return s[startIndex : startIndex+endIndex]
}
