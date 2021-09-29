package main

import (
	"fmt"
	"strings"
)

func convertIntArrToStr(a []uint) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), " []")
}

func CreatePhoneNumber(numbers [10]uint) string {
	return fmt.Sprintf("(%s) %s-%s",
		convertIntArrToStr(numbers[0:3]),
		convertIntArrToStr(numbers[3:6]),
		convertIntArrToStr(numbers[6:10]),
	)
}

func main() {
	fmt.Println("vim-go")
}
