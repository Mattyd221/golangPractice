package main

import (
	"github.com/Mattyd221/golangPractice/myUtils/fileUtils"
	"fmt"
)

func main() {
	// + = Static
	// ! = file pos
	// # = random Number
	// $ = spaces
	fmt.Println(fileUtils.LineDefinition("#2+1!0+testing!2$2+ha#3+ha"))
}
