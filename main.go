package main

import "fmt"

func main() {
	fmt.Println("Hello World 123")
	add(2, 3)
}

func add(a, b int) {
    fmt.Println(a + b)
}