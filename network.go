// network.go
package main

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"net/http"
	"os"
	"strings"
)

func processURL(url string) (string, string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while sending request:", err)
		return "", ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response:", err)
		return "", ""
	}
	bodyString := string(body)

	slug, fileName := extractSlugAndFileName(bodyString)
	if slug == "" || fileName == "" {
		fmt.Println("Error: Unable to extract slug or file name.")
		return "", ""
	}

	ext := extractExtension(fileName)
	newLink := fmt.Sprintf("https://qiwi.lol/%s.%s", slug, ext)
	fmt.Println("New Download Link:", newLink)

	return newLink, fileName
}

func extractSlugAndFileName(bodyString string) (string, string) {
	slug := extractBetween(bodyString, "\\\"slug\\\":\\\"", "\\\"")
	fileName := extractBetween(bodyString, "\\\"fileName\\\":\\\"", "\\\"")
	return slug, fileName
}

func extractExtension(fileName string) string {
	return fileName[strings.LastIndex(fileName, ".")+1:]
}

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

func downloadFile(url, fileName string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading the file:", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating the file:", err)
		return
	}
	defer out.Close()

	size := resp.ContentLength
	bar := pb.Full.Start64(size)
	barReader := bar.NewProxyReader(resp.Body)

	_, err = io.Copy(out, barReader)
	if err != nil {
		fmt.Println("Error while writing to the file:", err)
		return
	}

	bar.Finish()
	fmt.Printf("File successfully downloaded: '%s'\n", fileName)
}
