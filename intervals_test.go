package intervals

import (
	"strings"
	"testing"
	"time"
)

// comparator function for Integers
var cmpInt = func(x int, y int) int {
	if x == y {
		return 0
	} else if x > y {
		return 1
	}

	return -1
}

// comparator function for strings
var cmpStr = func(x string, y string) int {
	return strings.Compare(x, y) // compare strings lexicographically
}

// comparator function for dates
var cmpDate = func(x time.Time, y time.Time) int {
	if x.After(y) {
		return 1
	} else if x.Before(y) {
		return -1
	}

	return 0
}

func TestIntervalCreation(t *testing.T) {
	//------------------ Ints -------------------

	// Interval[int] [1,4]
	t1 := New(1, 4, cmpInt)
	if t1.x != 1 {
		t.Errorf("expected: 1, got: %d", t1.x)
	}
	if t1.y != 4 {
		t.Errorf("expected: 4, got: %d", t1.y)
	}

	// Interval[int] [4,1]
	t2 := New(4, 1, cmpInt)
	if t2.x != 1 {
		t.Errorf("expected: 1, got: %d", t2.x)
	}
	if t2.y != 4 {
		t.Errorf("expected: 4, got: %d", t2.y)
	}

	//---------------- Strings ------------------

	// Interval[string] ["abc, "defg"]
	t3 := New("abc", "defg", cmpStr)
	if t3.x != "abc" {
		t.Errorf("expected abc, got %s", t3.x)
	}
	if t3.y != "defg" {
		t.Errorf("expected defg, got %s", t3.y)
	}

	// Interval[string] ["defg", "abc"]
	t4 := New("defg", "abc", cmpStr)
	if t4.x != "abc" {
		t.Errorf("expected abc, got %s", t4.x)
	}
	if t4.y != "defg" {
		t.Errorf("expected defg, got %s", t4.y)
	}

	//---------------- Dates --------------------
	d1 := time.Now().Add(-(24 * time.Hour))
	d2 := time.Now()

	// Interva[time.Time] [time.Now() - 1 Day, time.Now()]
	t5 := New(d1, d2, cmpDate)
	if !t5.x.Equal(d1) {
		t.Errorf("expected time.Now() - 1 Day, got %v", t5.x)
	}
	if !t5.y.Equal(d2) {
		t.Errorf("expected time.Now(), got %v", t5.y)
	}

	// Interva[time.Time] [time.Now(), time.Now() - 1 Day]
	t6 := New(d2, d1, cmpDate)
	if !t6.x.Equal(d1) {
		t.Errorf("expected time.Now() - 1 Day, got %v", t6.x)
	}
	if !t6.y.Equal(d2) {
		t.Errorf("expected time.Now(), got %v", t6.y)
	}
}

func TestCommonRelations(t *testing.T) {
	i2 := New[int](5, 10, cmpInt)

	type test[T any] struct {
		input interval[T]
		want  Relation
	}

	tests := map[string]test[int]{
		"before": {
			input: New[int](1, 3, cmpInt),
			want:  Before,
		},
		"meets": {
			input: New[int](1, 5, cmpInt),
			want:  Meets,
		},
		"overlaps": {
			input: New[int](1, 7, cmpInt),
			want:  Overlaps,
		},
		"finishedBy": {
			input: New[int](1, 10, cmpInt),
			want:  FinishedBy,
		},
		"contains": {
			input: New[int](1, 15, cmpInt),
			want:  Contains,
		},
		"starts": {
			input: New[int](5, 7, cmpInt),
			want:  Starts,
		},
		"equal": {
			input: New[int](5, 10, cmpInt),
			want:  Equal,
		},
		"startedBy": {
			input: New[int](5, 15, cmpInt),
			want:  StartedBy,
		},
		"during": {
			input: New[int](7, 9, cmpInt),
			want:  During,
		},
		"finishes": {
			input: New[int](7, 10, cmpInt),
			want:  Finishes,
		},
		"overlappedBy": {
			input: New[int](8, 15, cmpInt),
			want:  OverlappedBy,
		},
		"metBy": {
			input: New[int](10, 15, cmpInt),
			want:  MetBy,
		},
		"after": {
			input: New[int](13, 17, cmpInt),
			want:  After,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := test.input.Relate(i2)
			if res != test.want {
				t.Fatalf("expected: %v, got: %v", test.want, res)
			}
		})
	}
}

func TestRelationsWithEmptyIntervals(t *testing.T) {
	type test[T any] struct {
		input [2]interval[T]
		want  Relation
	}

	tests := map[string]test[int]{
		"empty intervals overlaps": {
			input: [2]interval[int]{
				New[int](3, 3, cmpInt),
				New[int](3, 7, cmpInt),
			},
			want: Overlaps,
		},
		"empty intervals overlapped by": {
			input: [2]interval[int]{
				New[int](7, 7, cmpInt),
				New[int](3, 7, cmpInt),
			},
			want: OverlappedBy,
		},
		"empty intervals overllaped by 2": {
			input: [2]interval[int]{
				New[int](3, 7, cmpInt),
				New[int](3, 3, cmpInt),
			},
			want: OverlappedBy,
		},
		"empty intervals overlaps 2": {
			input: [2]interval[int]{
				New[int](3, 7, cmpInt),
				New[int](7, 7, cmpInt),
			},
			want: Overlaps,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := test.input[0].Relate(test.input[1])
			if res != test.want {
				t.Fatalf("expected: %v, got: %v", test.want, res)
			}
		})
	}
}
