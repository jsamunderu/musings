package main

import "fmt"

func match(str string) int {
	if len(str) == 0 {
		return 0
	}
	if str[0] == '(' {
		return match(str[1:]) + 1
	} else if str[0] == ')' {
		return match(str[1:]) - 1
	}
	return match(str[1:])
}

func main() {
	input := "))(("
	matched := match(input) == 0
	if matched {
		fmt.Println("matched")
	} else {
		fmt.Println("not match")
	}
}
