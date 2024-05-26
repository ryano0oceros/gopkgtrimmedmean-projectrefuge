# Go Package for Computing Trimmed Mean

## Intent of `sample-usage-branch`

The intent of this branch is to see how a sample project `myproject` can use the `gopkgtrimmedmean-projectrefuge` package. 

Benchmarking was performed on this branch and files committed. After moving `trimmedmean.go` to the `$GOPATH`, I simply `cd` into src/myproject and ran the following for the outputs in the `main` branch's `README.md`:

```bash
go run main.go

Rscript basemean.R
```
