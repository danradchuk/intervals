package intervals

import (
	"testing"
)

type intervalWithMax[C Comparator[T], T any] struct {
	i   Interval[C, T]
	max T
}

func TestIntervalTreeInsert(t *testing.T) {
	intervalsSorted := []intervalWithMax[IntComparator, int]{
		{
			i:   New[IntComparator, int](0, 3),
			max: 3,
		},
		{
			i:   New[IntComparator, int](5, 8),
			max: 10,
		},
		{
			i:   New[IntComparator, int](6, 10),
			max: 10,
		},
		{
			i:   New[IntComparator, int](8, 9),
			max: 23,
		},
		{
			i:   New[IntComparator, int](15, 23),
			max: 23,
		},
		{
			i:   New[IntComparator, int](16, 21),
			max: 30,
		},
		{
			i:   New[IntComparator, int](17, 19),
			max: 20,
		},
		{
			i:   New[IntComparator, int](19, 20),
			max: 20,
		},
		{
			i:   New[IntComparator, int](25, 30),
			max: 30,
		},
		{
			i:   New[IntComparator, int](26, 26),
			max: 26,
		},
	}

	tree := NewIntervalTree[IntComparator, int](New[IntComparator, int](16, 21))

	tree.Insert(New[IntComparator](8, 9))
	tree.Insert(New[IntComparator](25, 30))
	tree.Insert(New[IntComparator](5, 8))
	tree.Insert(New[IntComparator](15, 23))
	tree.Insert(New[IntComparator](17, 19))
	tree.Insert(New[IntComparator](0, 3))
	tree.Insert(New[IntComparator](6, 10))
	tree.Insert(New[IntComparator](19, 20))
	tree.Insert(New[IntComparator](26, 26))

	res := make([]intervalWithMax[IntComparator, int], 0)
	accumFunc := func(acc *[]intervalWithMax[IntComparator, int], i intervalWithMax[IntComparator, int]) {
		*acc = append(*acc, i)
	}
	traverseInOrder[IntComparator, int](tree.root, &res, accumFunc)

	for i, v := range res {
		if intervalsSorted[i].i.a != v.i.a {
			t.Errorf("LowEndpoint: expected %v, got %v", intervalsSorted[i].i.a, v.i.a)
		}

		if intervalsSorted[i].i.b != v.i.b {
			t.Errorf("LowEndpoint: expected %v, got %v", intervalsSorted[i].i.b, v.i.b)
		}

		if intervalsSorted[i].max != v.max {
			t.Errorf("Maximum in subtree: expected %v, got %v", intervalsSorted[i].max, v.max)
		}
	}
}

func traverseInOrder[C Comparator[T], T any](
	n *node[C, T],
	s *[]intervalWithMax[C, T],
	f func(*[]intervalWithMax[C, T], intervalWithMax[C, T]),
) {
	if n == nil {
		return
	}

	traverseInOrder(n.left, s, f)
	f(s, intervalWithMax[C, T]{n.val, n.max})
	traverseInOrder(n.right, s, f)
}

func TestSearchOverlaps(t *testing.T) {
	overlappedIntervals := []Interval[IntComparator, int]{
		{5, 8},
		{6, 10},
		{8, 9},
	}

	tree := NewIntervalTree[IntComparator, int](New[IntComparator, int](16, 21))

	tree.Insert(New[IntComparator, int](8, 9))
	tree.Insert(New[IntComparator, int](25, 30))
	tree.Insert(New[IntComparator, int](5, 8))
	tree.Insert(New[IntComparator, int](15, 23))
	tree.Insert(New[IntComparator, int](17, 19))
	tree.Insert(New[IntComparator, int](0, 3))
	tree.Insert(New[IntComparator, int](6, 10))
	tree.Insert(New[IntComparator, int](19, 20))
	tree.Insert(New[IntComparator, int](26, 26))

	s := tree.SearchOverlaps(New[IntComparator, int](4, 15))

	for i, v := range s {
		if overlappedIntervals[i].a != v.a {
			t.Errorf("LowEndpoint: expected %v, got %v", v.a, overlappedIntervals[i].a)
		}

		if overlappedIntervals[i].b != v.b {
			t.Errorf("HighEndpoint: expected %v, got %v", v.b, overlappedIntervals[i].b)
		}
	}
}
