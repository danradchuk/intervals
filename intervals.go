// Package intervals provides types and methods for working with closed ranges [a, b], where a and b can be any type that supports the <, >, and = operations.
package intervals

// Relation represents the relationship between interval x and interval y
type Relation int

const (
	// Added for exhaustiveness. Should never happen
	NoRel = iota

	/*
	   +---+
	   | x |
	   +---+
	        +---+
	        | y |
	        +---+
	*/
	Before

	/*
	   +---+
	   | x |
	   +---+
	       +---+
	       | y |
	       +---+
	*/
	Meets

	/*
	   +---+
	   | x |
	   +---+
	     +---+
	     | y |
	     +---+
	*/
	Overlaps

	/*
	   +-----+
	   |  x  |
	   +-----+
	     +---+
	     | y |
	     +---+
	*/

	FinishedBy

	/*
	   +-------+
	   |   x   |
	   +-------+
	     +---+
	     | y |
	     +---+
	*/
	Contains

	/*
	   +---+
	   | x |
	   +---+
	   +-------+
	   |   y   |
	   +-------+
	*/
	Starts

	/*
	   +---+
	   | x |
	   +---+
	   +---+
	   | y |
	   +---+
	*/
	Equal

	/*
	   +-----+
	   |  x  |
	   +-----+
	   +---+
	   | y |
	   +---+
	*/
	StartedBy

	/*
	     +---+
	     | x |
	     +---+
	   +-------+
	   |   y   |
	   +-------+
	*/
	During

	/*
	       +---+
	       | x |
	       +---+
	   +-------+
	   |   y   |
	   +-------+
	*/
	Finishes

	/*
	      +---+
	      | x |
	      +---+
	   +---+
	   | y |
	   +---+
	*/
	OverlappedBy

	/*
	       +---+
	       | x |
	       +---+
	   +---+
	   | y |
	   +---+
	*/
	MetBy

	/*
	         +---+
	         | x |
	         +---+
	   +---+
	   | y |
	   +---+
	*/
	After
)

// An Interval represents a closed range [a,b] of an arbitrary type T
type Interval[C Comparator[T], T any] struct {
	a T // low endpoint
	b T // high endpoint
}

// New creates a new Interval of type T
// If a > b, then elements of the interval are swapped
func New[C Comparator[T], T any](a, b T) Interval[C, T] {
	var c C
	cmpFunc := c.Compare

	if cmpFunc(a, b) == 1 {
		return Interval[C, T]{b, a}
	}

	return Interval[C, T]{a, b}
}

// LowEndpoint returns the low endpoint of the current interval (i.e. a from [a,b])
func (i Interval[C, T]) LowEndpoint() T {
	return i.a
}

// HighEndpoint returns the high endpoint of the current interval (i.e. b from [a,b])
func (i Interval[C, T]) HighEndpoint() T {
	return i.b
}

// Relate returns the relationship between two intervals of type T
func (i Interval[C, T]) Relate(other Interval[C, T]) Relation {
	var c C
	cmpFunc := c.Compare

	aa := cmpFunc(i.a, other.a)
	bb := cmpFunc(i.b, other.b)
	ba := cmpFunc(i.b, other.a)
	ab := cmpFunc(i.a, other.b)

	switch {
	case ba == -1:
		return Before
	case aa == -1 && ba == 0 && bb == -1:
		return Meets
	case ba == 0: // special case for zero intervals (i.e. (5,5))
		return Overlaps
	case aa == 1 && ab == 0 && bb == 1:
		return MetBy
	case ab == 0: // special case for zero intervals (i.e. (5,5))
		return OverlappedBy
	case ab == 1:
		return After
	case aa == -1 && bb == -1:
		return Overlaps
	case aa == -1 && bb == 0:
		return FinishedBy
	case aa == -1 && bb == 1:
		return Contains
	case aa == 0 && bb == -1:
		return Starts
	case aa == 0 && bb == 0:
		return Equal
	case aa == 0 && bb == 1:
		return StartedBy
	case aa == 1 && bb == -1:
		return During
	case aa == 1 && bb == 0:
		return Finishes
	case aa == 1 && bb == 1:
		return OverlappedBy
	default:
		return NoRel
	}
}

// IsEmpty reports whether the interval is empty
func (i Interval[C, T]) IsEmpty() bool {
	var c C
	cmpFunc := c.Compare

	return cmpFunc(i.a, i.b) == 0
}
