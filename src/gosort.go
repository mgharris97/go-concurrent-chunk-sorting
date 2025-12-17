package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
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
	// Range of values for random integer generation
	const MAX int = 1000
	const MIN int = 1

	// Check to see if user enters less than 3 args in terminal
	if len(os.Args) < 3 {
		fmt.Println("Usage: \n-r <number> \n-i <file path to .txt> \n-d <incoming directory of .txt files>")
		return
	}
	
	// Check to see if the os.Args is valid integer greater than 0
	n, err := strconv.Atoi(os.Args[2])
		if err != nil || n == 0 {
			fmt.Println("Error: argument must be a positive integer greater than 0")
			return
		}

		rand.Seed(time.Now().UnixNano())
		nums := make([]int, n)

		for i:=0, i<n, i++ {
			nums[i] = rand.Intn(MAX-MIN) + MIN
		}
	}
	



	//fmt.Println(n)
	//fmt.Println("input was" + os.Args[2])
	//fmt.Println(os.Args[1])

}
