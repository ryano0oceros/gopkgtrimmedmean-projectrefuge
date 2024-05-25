package trimmedmean

import (
	"sort"
)

// ComputeTrimmedMean calculates the trimmed mean of a slice of numbers.
// It supports both symmetric and asymmetric trimming.
// If only one trimming argument is provided, symmetric trimming is applied.
func ComputeTrimmedMean(data []float64, lowerTrim, upperTrim float64) float64 {
	if lowerTrim == 0 && upperTrim == 0 {
		return mean(data)
	}

	sort.Float64s(data)
	n := len(data)
	lowerCount := int(float64(n) * lowerTrim)
	upperCount := int(float64(n) * upperTrim)

	if lowerCount == upperCount {
		data = data[lowerCount : n-upperCount]
	} else {
		data = data[lowerCount : n-upperCount]
	}

	return mean(data)
}

// mean calculates the mean of a slice of numbers.
func mean(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}
