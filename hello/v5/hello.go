package main

import "fmt"

const PREFIX = "Hello,"

func Hello(name string) string {
	if name == "" {
		name = "Go"
	}
	return PREFIX + name
}

func main() {
	fmt.Println(Hello("World"))
}
