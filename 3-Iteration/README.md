# Iterations in Go - Learning Notes

## Key Concepts Learned

1. **Benchmarking**

- Benchmarking is a first class feature built into go
- Must start with `Benchmark`
- Makes use of `*testing.B` which gives us the amount of ideal iterations the function should go through to give a good result

2. **Types of loops**

- There's only for loop in go
- Any other loops like while, do while or until can be achieved solely with for