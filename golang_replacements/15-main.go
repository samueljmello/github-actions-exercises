package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	red = color.New(color.FgRed).SprintFunc()
)

func main() {
	fmt.Println(hello("GitHub"))
}

func hello(name string) (statement string) {
	return fmt.Sprintf("Hello, %s.", red(name))
}
