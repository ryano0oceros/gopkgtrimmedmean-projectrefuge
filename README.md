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

```go
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

To test the Go trimmed mean function against R's base mean function, you can use the following R script:

```R
x <- c(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
trimmedMeanR <- mean(x, trim = 0.05)
print(paste("Trimmed Mean in R:", trimmedMeanR))
```

Compare the output of this script with the output of your Go program to verify the correctness of the Go implementation.

## Running a Bootstrap Study

(Optional) To run a bootstrap study showing the value of the trimmed mean as a robust estimator of central tendency, refer to the `bootstrap_study.go` file in this repository. This file contains a sample implementation of a bootstrap study using the trimmed mean function.

