# Sample Project

## Improvements

- Update the data collecting mechanism from a file or any source (now it is hard coded in the main).
- Improve the data reading mechnism using stream reader, then large data files can be read.
- Use multiple go routines to process the data (specially with AggregateResult function under service) to improve the performance.