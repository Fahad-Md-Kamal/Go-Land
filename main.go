package main

import (
	"fmt"

	twosum "github.com/fahad-md-kamal/go-dsa-75/problems/twoSum"
	validparentheses "github.com/fahad-md-kamal/go-dsa-75/problems/valid-parentheses"
)


func main() {
	arr := []int{3, 4, 5, 6, 7}

	twoSum := twosum.IsValid(&arr, 7)
	fmt.Println(twoSum)

	fmt.Println("Enter Parenthesis for validatin check.")
	var input string
	fmt.Scanln(&input)
	ValidParenthesis := validparentheses.ValidParenthesis(&input)
	fmt.Println(ValidParenthesis)
}

