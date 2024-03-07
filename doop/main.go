package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	result := 0
	err := ""

	if len(input) != 3 {
		return
	}
	val1, error1 := strconv.Atoi(input[0])
	oper := input[1]
	val2, error2 := strconv.Atoi(input[2])

	if error1 != nil {
		return
	}

	if error2 != nil {
		return
	}

	if val1 >= 9223372036854775806 || val2 >= 9223372036854775806 {
		return
	}

	if val1 <= -9223372036854775806 || val2 <= -9223372036854775806 {
		return
	}

	switch oper {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "/":
		if val2 == 0 {
			err = "No division by 0"
		} else {
			result = val1 / val2
		}
	case "*":
		result = val1 * val2
	case "%":
		if val2 == 0 {
			err = "No modulo by 0"
		} else {
			result = val1 % val2
		}
	}

	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
