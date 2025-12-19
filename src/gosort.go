/*
	Matthew Harris
	241ADB166
	Graded Go Lab
	Dec. 20, 2025
*/

package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
)

func main() {

	// range of values for random integer generation
	const MAX int = 1000
	const MIN int = 1

	// check # of arguments entered
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("-r <number>")
		fmt.Println("-i <file path to .txt>")
		fmt.Println("-d <incoming directory of .txt files>")
		return
	}

	// convert input argument to an integer
	n, err := strconv.Atoi(os.Args[2])
	if err != nil || n < 10 {
		fmt.Println("Error: argument must be a positive integer at least 10")
		return
	}

	// create a slice of 'n' number of ints
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = rand.IntN(MAX-MIN+1) + MIN
		// test print to see if it works
		fmt.Printf("Value %d: %d\n", i, nums[i])
	}

	// chunking
	// number of chunks = ceil(sqrt(n))
	numOfChunks := int(math.Ceil(math.Sqrt(float64(n))))
	if numOfChunks < 4 {
		numOfChunks = 4
	}
	// chunck size should be roughly equal. Total number of ints / number of chunks = chunk size
	chunkSize := int(math.Ceil(float64(n) / float64(numOfChunks)))
	chunks := slices.Chunk(nums, chunkSize)

	// for testing purposes
	fmt.Printf("n = %d, numOfChunks = %d, chunkSize = %d\n", n, numOfChunks, chunkSize)

	i := 0
	for c := range chunks {
		fmt.Printf("Chunk %d size: %d\n", i, len(c))
		i++
	}

}

//fmt.Println(n)
//fmt.Println("input was" + os.Args[2])
//fmt.Println(os.Args[1])
