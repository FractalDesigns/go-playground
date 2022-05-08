package main

import (
	"go-playground/cmd"
)

func addition(a int, b int) int {
	return a + b
}

func main() {
	cmd.Execute()
}
