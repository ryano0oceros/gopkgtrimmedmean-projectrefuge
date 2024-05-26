package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"trimmedmean"
)

func readIntsFromFile(filename string) ([]int, error) {
	start := time.Now()

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ints []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("Time taken to read ints: %v\n", time.Since(start))

	return ints, nil
}

func readFloatsFromFile(filename string) ([]float64, error) {
	start := time.Now()

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var floats []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, f)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("Time taken to read floats: %v\n", time.Since(start))

	return floats, nil
}

func main() {
	data, err := readIntsFromFile("ints.txt")
	if err != nil {
		fmt.Println("Error reading ints:", err)
		return
	}

	// Convert []int to []float64
	floatData := make([]float64, len(data))
	for i, v := range data {
		floatData[i] = float64(v)
	}

	start := time.Now()
	trimmedMean := trimmedmean.ComputeTrimmedMean(floatData, 0.05, 0.05)
	fmt.Printf("Time taken to compute trimmed mean for 10,000 ints: %v\n", time.Since(start))

	fmt.Println("Trimmed Mean Ints:", trimmedMean)

	data2, err := readFloatsFromFile("floats.txt")
	if err != nil {
		fmt.Println("Error reading floats:", err)
		return
	}

	start2 := time.Now()
	trimmedMean2 := trimmedmean.ComputeTrimmedMean(data2, 0.05, 0.05)
	fmt.Printf("Time taken to compute trimmed mean for 100,000 floats: %v\n", time.Since(start2))

	fmt.Println("Trimmed Mean Floats:", trimmedMean2)
}
