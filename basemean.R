# Base mean function in R for trimmed mean comparison
# This function calculates the mean of a numeric vector, trimming a specified proportion of observations from each end.

calculateTrimmedMean <- function(x, trim = 0.05) {
  # Sort the vector to apply trimming
  x <- sort(x)
  
  # Calculate the number of observations to trim
  trimCount <- floor(trim * length(x))
  
  # Remove the specified proportion of observations from each end
  xTrimmed <- x[(trimCount + 1):(length(x) - trimCount)]
  
  # Calculate and return the mean of the trimmed observations
  mean(xTrimmed)
}

# Example usage
# x <- c(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
# trimmedMean <- calculateTrimmedMean(x, trim = 0.05)
# print(trimmedMean)
