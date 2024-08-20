package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	for range arr {
		fmt.Println(1)
	}
	fmt.Println("hello")
}
