# Go Slice Modification During Iteration

This repository demonstrates a common error in Go: modifying a slice while iterating over it using a range loop.  The provided `bug.go` file contains code that attempts to remove elements from a slice, which causes index out-of-bounds errors because the slice's length changes, causing index invalidation.

The `bugSolution.go` file presents a correct approach which avoids the problem and shows how to handle slice modification efficiently and safely.