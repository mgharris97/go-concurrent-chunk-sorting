/*
	Matthew Harris
	241ADB166
	Graded Go Lab
	Dec. 20, 2025
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"sync"
)

func main() {

	const MAX int = 1000
	const MIN int = 1

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("-r <number>")
		fmt.Println("-i <input.txt>")
		fmt.Println("-d <input_directory>")
		return
	}

	mode := os.Args[1]

	var nums []int

	switch mode {
	case "-r":
		n, err := strconv.Atoi(os.Args[2])
		if err != nil || n < 10 {
			fmt.Println("Error: argument must be a positive integer at least 10")
			return
		}

		nums = make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = rand.IntN(MAX-MIN+1) + MIN
		}

	case "-i":
		file, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Println("Error: cannot open input file")
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error: invalid integer in input file")
				return
			}
			nums = append(nums, val)
		}

	case "-d":
		files, err := filepath.Glob(filepath.Join(os.Args[2], "*.txt"))
		if err != nil || len(files) == 0 {
			fmt.Println("Error: no .txt files found")
			return
		}

		for _, f := range files {
			file, err := os.Open(f)
			if err != nil {
				fmt.Println("Error: cannot open file", f)
				return
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				val, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("Error: invalid integer in file", f)
					file.Close()
					return
				}
				nums = append(nums, val)
			}
			file.Close()
		}

	default:
		fmt.Println("Error: unknown mode. Must be -r, -i, or -d")
		return
	}

	// print unsorted nums
	fmt.Println("Original numbers (unsorted):")
	for _, v := range nums {
		fmt.Println(v)
	}
	fmt.Println()

	numOfChunks := int(math.Ceil(math.Sqrt(float64(len(nums)))))
	if numOfChunks < 4 {
		numOfChunks = 4
	}

	chunkSize := int(math.Ceil(float64(len(nums)) / float64(numOfChunks)))
	var chunks [][]int
	for c := range slices.Chunk(nums, chunkSize) {
		chunks = append(chunks, c)
	}

	// print chunks before sorting
	fmt.Println("Chunks before sorting:")
	for i, c := range chunks {
		fmt.Printf("Chunk %d:\n", i)
		for _, v := range c {
			fmt.Println(v)
		}
	}
	fmt.Println()

	var wg sync.WaitGroup

	for i := range chunks {
		wg.Add(1)
		go func(c []int) {
			defer wg.Done()
			sort.Ints(c)
		}(chunks[i])
	}

	wg.Wait()

	// print chunks after sorting
	fmt.Println("Chunks after sorting:")
	for i, c := range chunks {
		fmt.Printf("Chunk %d:\n", i)
		for _, v := range c {
			fmt.Println(v)
		}
	}
	fmt.Println()

	result := chunks[0]
	for i := 1; i < len(chunks); i++ {
		result = merge(result, chunks[i])
	}

	fmt.Println("Final merged sorted result:")
	for _, v := range result {
		fmt.Println(v)
	}

}

//fmt.Println(n)
//fmt.Println("input was" + os.Args[2])
//fmt.Println(os.Args[1])

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
