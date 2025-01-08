package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Check for correct arg count and open file for reading
	if len(os.Args) != 2 {
		fmt.Println("Correct usage go run . data.txt")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	// Init scanner, read line by line append to []int and sort
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	numbers := []int{}
	for scanner.Scan() {
		if n, err := strconv.Atoi(scanner.Text()); err == nil {
			numbers = append(numbers, n)
		}
	}
	if len(numbers) == 0 {
		fmt.Println("Incorrect data file")
		os.Exit(1)
	}
	sort.Ints(numbers)
	// Print calculations
	avrg, favg := average(numbers)
	medn := median(numbers)
	varian, deviat := variance(numbers, favg)
	fmt.Printf("Average: %v\n", avrg)
	fmt.Printf("Median: %v\n", medn)
	fmt.Printf("Variance: %v\n", varian)
	fmt.Printf("Standard Deviation: %v\n", deviat)
}

func average(numbers []int) (int, float64) {
	// Calculate sum of all numbers and divide by number count
	var sum float64
	for _, v := range numbers {
		sum += float64(v)
	}
	return int(math.Round(float64(sum / float64(len(numbers))))), float64(sum / float64(len(numbers)))

}
func median(numbers []int) int {
	// Check if even amount of numbers or odd and calculate median accordingly
	numbcount := len(numbers)
	if numbcount%2 == 0 {
		return int(math.Round((float64(numbers[numbcount/2]) + float64(numbers[numbcount/2-1]))/2))
	} else {
		return numbers[numbcount/2]
	}
}

func variance(numbers []int, avrg float64) (int, int) {
	// (x - avrg) squared
	var sum float64
	for _, v := range numbers {
		sum += math.Pow(float64(v)-avrg, 2)

	}
	varian := int(math.Round(float64(float64(sum) / float64(len(numbers)))))
	// Standard deviation is square root of variance
	deviat := int(math.Round(math.Sqrt((sum / float64(len(numbers))))))
	return varian, deviat
}
