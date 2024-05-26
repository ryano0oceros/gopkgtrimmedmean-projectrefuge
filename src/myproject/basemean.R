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

# Function to read numbers from a file
readNumbersFromFile <- function(filename) {
  as.numeric(readLines(filename))
}

# Benchmarking and calculating trimmed mean for ints.txt
start.time1 <- Sys.time()
ints <- readNumbersFromFile("ints.txt")
end.time1 <- Sys.time()
time.taken1 <- end.time1 - start.time1

start.time2 <- Sys.time()
trimmedMeanInts <- calculateTrimmedMean(ints, trim = 0.05)
end.time2 <- Sys.time()
time.taken2 <- end.time2 - start.time2
print(paste("Time taken for ints.txt: ", time.taken1))
print(paste("Time taken to compute trimmed mean for 10,000 ints: ", time.taken2))
print(paste("Trimmed Mean Ints: ", trimmedMeanInts))

# Benchmarking and calculating trimmed mean for floats.txt
start.time3 <- Sys.time()
floats <- readNumbersFromFile("floats.txt")
end.time3 <- Sys.time()
time.taken3 <- end.time3 - start.time3

start.time4 <- Sys.time()
trimmedMeanFloats <- calculateTrimmedMean(floats, trim = 0.05)
end.time4 <- Sys.time()
time.taken4 <- end.time4 - start.time4
print(paste("Time taken for floats.txt: ", time.taken3))
print(paste("Time taken to compute trimmed mean for 100,000 floats: ", time.taken4))
print(paste("Trimmed Mean Floats: ", trimmedMeanFloats))