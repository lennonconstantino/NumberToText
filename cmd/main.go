package main

import (
	parseNumberToTest "NumberToText/src/parseNumberToText"
	"fmt"
)

func main() {
	fmt.Println("Starting...")

	listIntegers := []int64{
		1,
		10,
		11,
		12,
		19,
		21,
		23,
		33,
		50,
		60,
		70,
		100,
		101,
		123,
		247,
		503,
		1234,
		12345,
		10000,
		10001,
		33300,
		99900,
		100000,
		100001,
		123123,
		224501,
		7890554,
		13456787,
		100000001,
		123123123,
		100000002,
		1_231_231_231_23,
		10_129_991_234_44,
		123_123_123_123_123,
		999_666_234_567_830_01,
		999_666_234_567_890_01,
		99999999999999999,
		100000000000000000,
	}

	for i := range listIntegers {
		bufferNumberToText, err := parseNumberToTest.NumberToText(listIntegers[i])
		if err != nil {
			fmt.Println("Business error: ", err)
			continue
		}

		fmt.Println(bufferNumberToText)
	}
}
