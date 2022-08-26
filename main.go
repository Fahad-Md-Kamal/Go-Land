package main

import (
	"fmt"

	twosum "github.com/fahad-md-kamal/go-dsa-75/problems/twoSum"
)


func main() {
	arr := []int{3, 4, 5, 6, 7}

	twoSum := twosum.IsValid(&arr, 7)
	fmt.Println(twoSum)
}

