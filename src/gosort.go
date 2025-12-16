package main

import (
	"fmt"
	"os"
	"strconv"
)

func isInt(input string) bool {
	value := false
	n, err := strconv.Atoi(input)

	if err == nil {
		value = true
	}

	if n < 0 {
		value = false
	}
	return value
	// using strconv.Atoi() to convert string to int
	// converted int goes in 'n' and err is set to nil
	// if conversion fails, then n gets default value 0 and error object (info) goes into err
}

func main() {

	if !isInt(os.Args[2]) {
		fmt.Println("Error: argument must be a non-negative integer")
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("Usage: \n-r <number> \n-i <file path to .txt> \n-d <incoming directory of .txt files>")
		return
	}

	//fmt.Println(n)
	//fmt.Println("input was" + os.Args[2])
	//fmt.Println(os.Args[1])

}
