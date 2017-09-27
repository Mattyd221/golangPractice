package main

import (
	"bufio"
	"fmt"
	"time"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	baseTime := time.Now()
	readLines("F:\\Go Projects\\output.txt")
	fmt.Println(time.Since(baseTime))
	baseTime = time.Now()
	readAll("F:\\Go Projects\\output.txt")
	fmt.Println(time.Since(baseTime))
}


func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


func readAll(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	content, err2 := ioutil.ReadAll(reader)

	return strings.Split(string(content),"\n"), err2
}