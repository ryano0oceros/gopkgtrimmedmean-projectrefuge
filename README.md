# Go Package for Computing Trimmed Mean

## Getting Started

To use the Go package for computing a trimmed mean, first ensure you have Go installed on your system. Then, create a new directory for your project and initialize a new Go module:

```bash
mkdir myproject
cd myproject
go mod init myproject
```

Next, create a file named `trimmedmean.go` in your project directory. You can find the implementation details for the trimmed mean function in the `trimmedmean.go` file section of this repository.

## Using the Trimmed Mean Function

The trimmed mean function can compute both symmetric and asymmetric trimming based on the provided arguments. Here is an example of how to use it:

```bash
package main

import (
    "fmt"
    "myproject/trimmedmean"
)

func main() {
    data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    trimmedMean := trimmedmean.ComputeTrimmedMean(data, 0.1, 0.1)
    fmt.Println("Trimmed Mean:", trimmedMean)
}
```

## Testing Against R's Base Mean Function

See `sample-usage-branch`

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
