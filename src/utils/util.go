package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputInt64() (int64, error) {
	fmt.Println("Enter an integer")
	userInput := bufio.NewReader(os.Stdin)
	userVal, err := userInput.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input := strings.TrimSpace(userVal)
	intVal64, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		fmt.Printf("Error to convert: Number is larger")
		return 0, err
	}

	fmt.Printf("You entered: %d\n", intVal64)
	return intVal64, nil
}

func NumberToText(number int64) string {
	result := fmt.Sprintf("%d", number)
	return result
}
