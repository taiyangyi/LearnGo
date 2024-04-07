package main

import "fmt"

const PREFIX = "Hello,"

func Hello(name string) string {
	return PREFIX + name
}

func main() {

	fmt.Println(Hello("Go"))

}
