package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	regex, err := regexp.Compile(`\/var\/www.*?\s`)
	if err != nil {
		panic(err.Error())
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if regex.MatchString(line) {
			wg.Add(1)
			path := strings.TrimSuffix(regex.FindString(line), " ")
			go checkPathExists(path)
		}
	}
	wg.Wait()
}

func checkPathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Path " + path + " does not exists!")
			return false
		}
	}
	defer wg.Done()
	return true
}
