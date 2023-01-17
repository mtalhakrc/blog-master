package main

import "fmt"

func main() {
	son(ilk())
}

func ilk() string {
	fmt.Println("ilk")
	return ""
}
func son(str string) {
	fmt.Println("son")
}
