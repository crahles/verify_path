package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	regex, err := regexp.Compile(`\/var\/www.*?\s`)
	if err != nil {
		panic(err.Error())
	}

	matchingLines := readMatchingLines(os.Stdin, regex)
	outputChannel := make(chan string)

	for _, line := range matchingLines {
		path := strings.TrimSuffix(regex.FindString(line), " ")
		go checkPathExists(outputChannel, path)
	}

	for i := 0; i <= len(matchingLines); i++ {
		if msg := <-outputChannel; msg != "" {
			fmt.Println(msg)
		}
	}
}

func readMatchingLines(input io.Reader, regex *regexp.Regexp) []string {
	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if regex.MatchString(line) {
			lines = append(lines, line)
		}
	}
	return lines
}

func checkPathExists(cs chan string, path string) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			cs <- "Path " + path + " does not exists!"
		}
	}
	cs <- ""
}
