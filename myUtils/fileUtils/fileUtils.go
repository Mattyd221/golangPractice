package fileUtils

import (
	"os"
	"bufio"
	"io/ioutil"
	"strings"
	"bytes"
	"github.com/Mattyd221/golangPractice/myUtils/stringUtils"
	"strconv"
	"math/rand"
	"time"
	"fmt"
)

func ReadLines(path string) ([]string, error) {
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


func ReadAll(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	content, err2 := ioutil.ReadAll(reader)

	return strings.Split(string(content),"\r\n"), err2
}

func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  writer := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(writer, line)
  }
  return writer.Flush()
}

func LineDefinition(definitionString string) string {
	// + = Static
	// ! = file pos
	// # = random Number
	// $ = spaces
	inArr := strings.Split(definitionString,"")
	var buffer bytes.Buffer
	var currentToken = ""
	var tokenString = ""
	rand.Seed(int64(time.Now().Second()))
	for _, x := range inArr {
		switch x {
		case "+":
			buffer.WriteString(calcString(currentToken, tokenString))
			currentToken = "+"
			tokenString = ""
		case "!":
			buffer.WriteString(calcString(currentToken, tokenString))
			currentToken = "!"
			tokenString = ""
		case "#":
			buffer.WriteString(calcString(currentToken, tokenString))
			currentToken = "#"
			tokenString = ""
		case "$":
			buffer.WriteString(calcString(currentToken, tokenString))
			currentToken = "$"
			tokenString = ""
		default:
			tokenString += x
		}

	}
	buffer.WriteString(calcString(currentToken, tokenString))
	return buffer.String()
}

func calcString(tokenChar string, tokenString string) string {
	switch tokenChar {
	case "+":
		return tokenString
	case "!":
		return "[" + tokenString + "]"
	case "#":
		ans, _ := strconv.Atoi(tokenString)
		return calcNum(ans)
	case "$":
		ans, _ := strconv.Atoi(tokenString)
		return stringUtils.RightPad(""," ",ans)
	default:
		return ""
	}

}

func calcNum(numLen int) string {
	var randNum int
	min, _ := strconv.Atoi(stringUtils.RightPad("1","0",numLen - 1))
	max, _ := strconv.Atoi(stringUtils.RightPad("9","9",numLen - 1))
	randNum = int(float64(max-min+1)*rand.Float64()+float64(min))
	return strconv.Itoa(randNum)
}
