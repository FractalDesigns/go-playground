package main

import (
	"github.com/afritfr/go-playground/cmd"
)

func addition(a int, b int) int {
	return a + b
}

func main() {
	cmd.Execute()
}
