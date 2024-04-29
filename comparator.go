package intervals

import (
	"strings"
	"time"
)

// Comparator[T any] is an interface used for implementing custom comparators for a specific type T
type Comparator[T any] interface {
	~struct{}
	Compare(a, b T) int
}

// IntComparator is a default comparator for int type
type IntComparator struct{}

// Compare returns 1 if a > b, 0 if  a == b, -1 otherwise
func (IntComparator) Compare(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}

	return -1
}

// StringComparator is a default comparator for string type
type StringComparator struct{}

// Compare strings lexicographically
func (StringComparator) Compare(a, b string) int {
	return strings.Compare(a, b)
}

// TimeComparator is a default comparator for time.Time type
type TimeComparator struct{}

// Compare returns 1 if a is after b, -1 if a before b; 0 if a represents the same time as b
func (TimeComparator) Compare(a, b time.Time) int {
	if a.After(b) {
		return 1
	} else if a.Before(b) {
		return -1
	}

	return 0
}
