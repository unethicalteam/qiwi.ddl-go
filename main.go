// main.go
package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
    clear()

	url := getURLFromArgsOrPrompt()
	if !isValidQiwiURL(url) {
		return
	}
    newLink, fileName := processURL(url)
		if newLink == "" {
		return
	}

	handleDownloadDecision(newLink, fileName)
}

func getURLFromArgsOrPrompt() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter download link: ")
	url, _ := reader.ReadString('\n')
	return strings.TrimSpace(url)
}

func isValidQiwiURL(url string) bool {
	if strings.HasPrefix(url, "https://qiwi.gg/folder") {
		fmt.Println("Error: Folder handling isn't implemented at the moment.")
		return false
	}  else if !strings.HasPrefix(url, "https://qiwi.gg/file") {
		fmt.Println("Error: Invalid URL. Only qiwi.gg/file URLs are supported.")
		return false
	}
	return true
}

func confirmDownload() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to download the file? (y/n): ")
	response, _ := reader.ReadString('\n')
	return strings.HasPrefix(strings.TrimSpace(strings.ToLower(response)), "y")
}

func handleDownloadDecision(newLink, fileName string) {
	if confirmDownload() {
		downloadFile(newLink, fileName)
	} else {
		err := clipboard.WriteAll(newLink)
		if err != nil {
			fmt.Println("Failed to copy to clipboard:", err)
		} else {
			fmt.Println("Download link copied to clipboard.")
		}
	}
}
