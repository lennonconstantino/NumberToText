package main

import (
	parseNumberToTest "NumberToText/src/parseNumberToText"
	"fmt"
)

func main() {
	fmt.Println("Starting...")

	fmt.Println(parseNumberToTest.NumberToText(1))
	fmt.Println(parseNumberToTest.NumberToText(12))
	fmt.Println(parseNumberToTest.NumberToText(123))
	fmt.Println(parseNumberToTest.NumberToText(123123))
	fmt.Println(parseNumberToTest.NumberToText(123123123))
	fmt.Println(parseNumberToTest.NumberToText(123123123123))
	fmt.Println(parseNumberToTest.NumberToText(123123123123123))

}
