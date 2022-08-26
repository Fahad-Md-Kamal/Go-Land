package arr

import "fmt"


func arrFunc() {

	var stack []string;
	
	stack = append(stack, "A")
	stack = append(stack, "B")
	stack = append(stack, "C")
	stack = append(stack, "D")
	
	fmt.Println(stack)
}