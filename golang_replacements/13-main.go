package main

import "fmt"

func main() {
	fmt.Println(hello("GitHub"))
}

func hello(name string) (statement string) {
	return fmt.Sprintf("Hello, %s.", name)
}

func badFunc() {
	fmt.Println("I'm not going to get used.")
}
