package trimmedmean

import (
	"testing"
)

// TestComputeTrimmedMeanIntegers tests the ComputeTrimmedMean function with a sample of integers.
func TestComputeTrimmedMeanIntegers(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	expectedTrimmedMean := 50.5 // Expected result with symmetric trimming of 0.05
	trimmedMean := ComputeTrimmedMean(data, 0.05, 0.05)

	if trimmedMean != expectedTrimmedMean {
		t.Errorf("ComputeTrimmedMean with integers = %f; want %f", trimmedMean, expectedTrimmedMean)
	}
}

// TestComputeTrimmedMeanFloats tests the ComputeTrimmedMean function with a sample of floats.
func TestComputeTrimmedMeanFloats(t *testing.T) {
	data := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.2, 12.3, 13.4, 14.5, 15.6, 16.7, 17.8, 18.9, 19.1, 20.2, 21.3, 22.4, 23.5, 24.6, 25.7, 26.8, 27.9, 28.1, 29.2, 30.3, 31.4, 32.5, 33.6, 34.7, 35.8, 36.9, 37.1, 38.2, 39.3, 40.4, 41.5, 42.6, 43.7, 44.8, 45.9, 46.1, 47.2, 48.3, 49.4, 50.5, 51.6, 52.7, 53.8, 54.9, 55.1, 56.2, 57.3, 58.4, 59.5, 60.6, 61.7, 62.8, 63.9, 64.1, 65.2, 66.3, 67.4, 68.5, 69.6, 70.7, 71.8, 72.9, 73.1, 74.2, 75.3, 76.4, 77.5, 78.6, 79.7, 80.8, 81.9, 82.1, 83.2, 84.3, 85.4, 86.5, 87.6, 88.7, 89.8, 90.9, 91.1, 92.2, 93.3, 94.4, 95.5, 96.6, 97.7, 98.8, 99.9, 100.1}
	expectedTrimmedMean := 50.6 // Expected result with symmetric trimming of 0.05
	trimmedMean := ComputeTrimmedMean(data, 0.05, 0.05)

	if trimmedMean != expectedTrimmedMean {
		t.Errorf("ComputeTrimmedMean with floats = %f; want %f", trimmedMean, expectedTrimmedMean)
	}
}
