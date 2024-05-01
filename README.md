# Intervals

Intervals is a tiny library for working with intervals inspired by a Haskell library, [Rampart](https://hackage.haskell.org/package/rampart-2.0.0.7).

# Install

`go get github.com/danradchuk/intervals`

# Usage

If you want to know how interval `i1` relates to interval `i2`, you can use `Relate` method:

```go
package main

import (
	"github.com/danradchuk/intervals"
)

func main() {
	i1 := intervals.New[intervals.IntComparator, int](1, 10)
	i2 := intervals.New[intervals.IntComparator, int](5, 8)

	i1.Relate(i2) // Contains
}
```

If you have an array of intervals and wish to determine if a specific interval overlaps with any of them, you can use the `intervals.IntervalTree[C,T]` data structure:

```go
package main

import "github.com/danradchuk/intervals"

func main() {
	rootKey := intervals.New[intervals.IntComparator, int](16, 21)
	tree := intervals.NewIntervalTree[intervals.IntComparator, int](rootKey)

	tree.Insert(intervals.New[intervals.IntComparator](8, 9))
	tree.Insert(intervals.New[intervals.IntComparator](25, 30))
	tree.Insert(intervals.New[intervals.IntComparator](5, 8))
	tree.Insert(intervals.New[intervals.IntComparator](15, 23))
	tree.Insert(intervals.New[intervals.IntComparator](17, 19))
	tree.Insert(intervals.New[intervals.IntComparator](0, 3))
	tree.Insert(intervals.New[intervals.IntComparator](6, 10))
	tree.Insert(intervals.New[intervals.IntComparator](19, 20))
	tree.Insert(intervals.New[intervals.IntComparator](26, 26))

	tree.SearchOverlaps(intervals.New[intervals.IntComparator, int](4, 15)) // returns a slice with intervals [5,8], [6,10], and [8,9]
}
```
