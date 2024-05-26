# Go Package for Computing Trimmed Mean

## Getting Started

To use the Go package for computing a trimmed mean, first ensure you have Go installed on your system. Then, create a new directory for your project and initialize a new Go module:

```bash
mkdir myproject
cd myproject
go mod init myproject
```

Additionally make sure to add the `trimmedmean.go` file into your $GOPATH. OS-specific instructions on how to do that can be found [here](https://go.dev/wiki/SettingGOPATH).

## Using the Trimmed Mean Function

The trimmed mean function can compute both symmetric and asymmetric trimming based on the provided arguments. Here is a simple example of how to use it:

```go
package main

import (
    "fmt"
    "trimmedmean"
)

func main() {
    data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    trimmedMean := trimmedmean.ComputeTrimmedMean(data, 0.1, 0.1)
    fmt.Println("Trimmed Mean:", trimmedMean)
}
```

## Testing Against R's Base Mean Function

See [`sample-usage-branch`](https://github.com/ryano0oceros/gopkgtrimmedmean-projectrefuge/blob/sample-usage-branch/README.md) for all the code that was used, and additional write-up.

### Go Results

```bash
Time taken to read ints: 46.274291ms
Time taken to compute trimmed mean for 10,000 ints: 1.08025ms
Trimmed Mean Ints: 4.641250483401553e+18
Time taken to read floats: 17.362417ms
Time taken to compute trimmed mean for 100,000 floats: 10.651042ms
Trimmed Mean Floats: 0.4998520822609632
```

### R Results

```bash
[1] "Time taken for ints.txt:  0.00249600410461426"
[1] "Time taken to compute trimmed mean for 10,000 ints:  0.000615119934082031"
[1] "Trimmed Mean Ints:  4641250483401547776"
[1] "Time taken for floats.txt:  0.0341808795928955"
[1] "Time taken to compute trimmed mean for 100,000 floats:  0.0104949474334717"
[1] "Trimmed Mean Floats:  0.499852082260964"
```

### Summary

Output of Go program is almost identical (each to 14 decimal places). Surprisingly, overall with these samples R outperformed Go. Part of this, at least for ints, is that I set up my package to take a slice of floats and for the ints I had to cast them to floats beforehand adding compute overhead. For the case of 100,000 floats the metrics are almost identical with Go reading values from floats.txt slightly faster, but R slightly outperforming Go in computing the trimmed mean. 

## Appendix

### `main.go` Code Used for Benchmarking

```go
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
```
