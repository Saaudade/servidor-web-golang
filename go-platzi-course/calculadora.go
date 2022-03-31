package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(input string, operator string) int {
	inputClean := strings.Split(input, operator)
	operator_1 := parsear(inputClean[0])
	operator_2 := parsear(inputClean[1])
	result := 0
	switch operator {
	case "+":
		result = operator_1 + operator_2
	case "-":
		result = operator_1 - operator_2
	case "*":
		result = operator_1 * operator_2
	case "/":
		result = operator_1 / operator_2
	default:
		fmt.Printf("Operator %s is not available\n", operator)
	}
	return result
}

func parsear(input string) int {
	operator, _ := strconv.Atoi(input)
	return operator
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	input := readInput()
	operator := readInput()

	calculate := calc{}
	result := calculate.operate(input, operator)

	fmt.Println(result)
}
