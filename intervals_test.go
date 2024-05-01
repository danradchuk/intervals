package intervals

import (
	"testing"
	"time"
)

func TestIntervalCreation(t *testing.T) {
	//------------------ Ints -------------------

	// Interval[int] [1,4]
	t1 := New[IntComparator, int](1, 4)
	if t1.a != 1 {
		t.Errorf("expected: 1, got: %d", t1.a)
	}
	if t1.b != 4 {
		t.Errorf("expected: 4, got: %d", t1.b)
	}

	// Interval[int] [4,1]
	t2 := New[IntComparator, int](4, 1)
	if t2.a != 1 {
		t.Errorf("expected: 1, got: %d", t2.a)
	}
	if t2.b != 4 {
		t.Errorf("expected: 4, got: %d", t2.b)
	}

	//---------------- Strings ------------------

	// Interval[string] ["abc, "defg"]
	t3 := New[StringComparator, string]("abc", "defg")
	if t3.a != "abc" {
		t.Errorf("expected abc, got %s", t3.a)
	}
	if t3.b != "defg" {
		t.Errorf("expected defg, got %s", t3.b)
	}

	// Interval[string] ["defg", "abc"]
	t4 := New[StringComparator, string]("defg", "abc")
	if t4.a != "abc" {
		t.Errorf("expected abc, got %s", t4.a)
	}
	if t4.b != "defg" {
		t.Errorf("expected defg, got %s", t4.b)
	}

	//---------------- Dates --------------------
	d1 := time.Now().Add(-(24 * time.Hour))
	d2 := time.Now()

	// Interval[time.Time] [time.Now() - 1 Day, time.Now()]
	t5 := New[TimeComparator, time.Time](d1, d2)
	if !t5.a.Equal(d1) {
		t.Errorf("expected time.Now() - 1 Day, got %v", t5.a)
	}
	if !t5.b.Equal(d2) {
		t.Errorf("expected time.Now(), got %v", t5.b)
	}

	// Interval[time.Time] [time.Now(), time.Now() - 1 Day]
	t6 := New[TimeComparator, time.Time](d2, d1)
	if !t6.a.Equal(d1) {
		t.Errorf("expected time.Now() - 1 Day, got %v", t6.a)
	}
	if !t6.b.Equal(d2) {
		t.Errorf("expected time.Now(), got %v", t6.b)
	}
}

func TestCommonRelations(t *testing.T) {
	i2 := New[IntComparator, int](5, 10)

	type test[C Comparator[T], T any] struct {
		input Interval[C, T]
		want  Relation
	}

	tests := map[string]test[IntComparator, int]{
		"before": {
			input: New[IntComparator, int](1, 3),
			want:  Before,
		},
		"meets": {
			input: New[IntComparator, int](1, 5),
			want:  Meets,
		},
		"overlaps": {
			input: New[IntComparator, int](1, 7),
			want:  Overlaps,
		},
		"finishedBy": {
			input: New[IntComparator, int](1, 10),
			want:  FinishedBy,
		},
		"contains": {
			input: New[IntComparator, int](1, 15),
			want:  Contains,
		},
		"starts": {
			input: New[IntComparator, int](5, 7),
			want:  Starts,
		},
		"equal": {
			input: New[IntComparator, int](5, 10),
			want:  Equal,
		},
		"startedBy": {
			input: New[IntComparator, int](5, 15),
			want:  StartedBy,
		},
		"during": {
			input: New[IntComparator, int](7, 9),
			want:  During,
		},
		"finishes": {
			input: New[IntComparator, int](7, 10),
			want:  Finishes,
		},
		"overlappedBy": {
			input: New[IntComparator, int](8, 15),
			want:  OverlappedBy,
		},
		"metBy": {
			input: New[IntComparator, int](10, 15),
			want:  MetBy,
		},
		"after": {
			input: New[IntComparator, int](13, 17),
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
	type test[C Comparator[T], T any] struct {
		input [2]Interval[C, T]
		want  Relation
	}

	tests := map[string]test[IntComparator, int]{
		"empty intervals overlaps": {
			input: [2]Interval[IntComparator, int]{
				New[IntComparator, int](3, 3),
				New[IntComparator, int](3, 7),
			},
			want: Overlaps,
		},
		"empty intervals overlapped by": {
			input: [2]Interval[IntComparator, int]{
				New[IntComparator, int](7, 7),
				New[IntComparator, int](3, 7),
			},
			want: OverlappedBy,
		},
		"empty intervals overllaped by 2": {
			input: [2]Interval[IntComparator, int]{
				New[IntComparator, int](3, 7),
				New[IntComparator, int](3, 3),
			},
			want: OverlappedBy,
		},
		"empty intervals overlaps 2": {
			input: [2]Interval[IntComparator, int]{
				New[IntComparator, int](3, 7),
				New[IntComparator, int](7, 7),
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
