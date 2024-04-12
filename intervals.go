package intervals

type Relation int

const (
	_ = iota

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

// Interval represents a range (i.e. [x,y])
type interval[T any] struct {
	x       T              // low endpoint
	y       T              // high endpoint
	compare func(T, T) int // a comparison function returns 1 if the first argument is bigger than the second one; 0 if they are equal; -1 otherwise
}

// Returns the low endpoint of the interval
func (i interval[T]) LowEndpoint() T {
	return i.x
}

// Returns the high endpoint of the interval
func (i interval[T]) HighEndpoint() T {
	return i.y
}

func (i interval[T]) Relate(other interval[T]) Relation {
	xx := i.compare(i.x, other.x)
	yy := i.compare(i.y, other.y)
	yx := i.compare(i.y, other.x)
	xy := i.compare(i.x, other.y)

	switch {
	case yx == -1:
		return Before
	case xx == -1 && yx == 0 && yy == -1:
		return Meets
	case yx == 0: // special case for zero intervals (i.e. (5,5))
		return Overlaps
	case xx == 1 && xy == 0 && yy == 1:
		return MetBy
	case xy == 0: // special case for zero intervals (i.e. (5,5))
		return OverlappedBy
	case xy == 1:
		return After
	case xx == -1 && yy == -1:
		return Overlaps
	case xx == -1 && yy == 0:
		return FinishedBy
	case xx == -1 && yy == 1:
		return Contains
	case xx == 0 && yy == -1:
		return Starts
	case xx == 0 && yy == 0:
		return Equal
	case xx == 0 && yy == 1:
		return StartedBy
	case xx == 1 && yy == -1:
		return During
	case xx == 1 && yy == 0:
		return Finishes
	case xx == 1 && yy == 1:
		return OverlappedBy
	default:
		return 0
	}
}

// Check if the interval is empty (e.g. (5,5))
func (i interval[T]) IsEmpty() bool {
	return i.compare(i.x, i.y) == 0
}

func New[T any](x T, y T, cmpFunc func(T, T) int) interval[T] {
	if cmpFunc(x, y) == 1 {
		return interval[T]{y, x, cmpFunc}
	}

	return interval[T]{x, y, cmpFunc}
}
