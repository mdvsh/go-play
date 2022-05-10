package main

import "fmt"

func Hello(name string) string {
	return "hi " + name
}

func main() {
	// fmt.Println("test driven dev...")
	fmt.Println(Hello("madhav"))
}
