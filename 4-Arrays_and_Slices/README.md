# Arrays and Slices in Go

## Key Concepts

### Arrays
- Arrays have a fixed length
- Array type is specified with both the size and type: `[5]int`
- Arrays are value types - when passed to functions, a copy is made
- Length is part of the array's type, so `[5]int` and `[4]int` are different types

### Slices
- Slices are more flexible than arrays as they can grow/shrink
- Slice type is specified without size: `[]int`
- Slices are reference types - they point to an underlying array
- `append` function is used to add elements to a slice
- Can create slices from arrays using `array[low:high]` syntax

### Range
- `range` keyword lets you iterate over arrays or slices
- Returns both index and value: `for i, v := range numbers`
- Can ignore index using underscore: `for _, v := range numbers`

### Variadic Functions
- Functions can accept variable number of arguments using `...`
- Example: `func SumAll(slices ...[]int)`
- Allows passing multiple slices as separate arguments
