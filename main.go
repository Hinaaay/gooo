package main

import "fmt"

func main() {
	var a, b int
	var c string
	fmt.Printf("input:")
	fmt.Scan(&a, &c, &b)
	switch c {
	case "+":
		fmt.Println(a, c, b, "=", a+b)
	case "-":
		fmt.Println(a, c, b, "=", a-b)
	case "*":
		fmt.Println(a, c, b, "*", a*b)
	case "/":
		fmt.Println(a, c, b, "/", a/b)
	default:
		fmt.Println("error")
	}
}
