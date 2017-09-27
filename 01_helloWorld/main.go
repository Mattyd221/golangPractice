package main

import (
	"github.com/Mattyd221/golangPractice/myUtils"
	"fmt"
)

func main() {

	fmt.Println(myUtils.RightPadToLength("0123456789"," ",16))
	fmt.Println(myUtils.LeftPadToLength("0123456789"," ",16))
}
