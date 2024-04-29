package intervals

// TODO replace BST with RBT

// Tree represents a standard Binary search tree with additional field max in its nodes
// max is the biggest HighEndpoint of an interval (i.e. b from [a,b]) in a subtree
type Tree[C Comparator[T], T any] struct {
	root *node[C, T]
}

type node[C Comparator[T], T any] struct {
	left   *node[C, T]
	val    Interval[C, T]
	right  *node[C, T]
	max    T
	parent *node[C, T]
}

// NewIntervalTree creates a new IntervalTree of type T
func NewIntervalTree[C Comparator[T], T any](i Interval[C, T]) *Tree[C, T] {
	return &Tree[C, T]{
		root: &node[C, T]{
			left:   nil,
			val:    i,
			right:  nil,
			max:    i.HighEndpoint(),
			parent: nil,
		},
	}
}

// Insert adds an interval to the tree.
// Ignore duplicates.
func (t *Tree[C, T]) Insert(i Interval[C, T]) {
	insert(&t.root, i, nil)
}

func insert[C Comparator[T], T any](n **node[C, T], i Interval[C, T], p *node[C, T]) {
	if *n == nil {
		*n = &node[C, T]{nil, i, nil, i.HighEndpoint(), p}
		return
	}

	var c C
	cmpFunc := c.Compare

	tmp := *n
	val := tmp.val
	cmpX := cmpFunc(i.a, val.a)
	cmpY := cmpFunc(i.b, val.b)

	// TODO deals with duplicates
	// compare low endpoints
	if cmpX == -1 {
		insert(&(tmp.left), i, tmp)
	} else if cmpX == 1 {
		insert(&(tmp.right), i, tmp)
	} else {
		// compare high endpoints
		if cmpY == -1 {
			insert(&(tmp.left), i, tmp)
		} else if cmpY == 1 {
			insert(&(tmp.right), i, tmp)
		} else {
			return
		}
	}

	if tmp.left == nil && tmp.right == nil { // no children
		tmp.max = tmp.val.HighEndpoint()
	} else if tmp.right == nil { // only left child
		tmp.max = maxEndpoint(cmpFunc, tmp.val.HighEndpoint(), tmp.left.val.HighEndpoint())
	} else if tmp.left == nil { // only right child
		tmp.max = maxEndpoint(cmpFunc, tmp.val.HighEndpoint(), tmp.right.val.HighEndpoint())
	} else { // both children
		tmp.max = maxEndpoint(cmpFunc, tmp.val.HighEndpoint(), tmp.left.val.HighEndpoint(), tmp.right.val.HighEndpoint())
	}
}

func maxEndpoint[T any](f func(T, T) int, args ...T) T {
	res := args[0]
	for _, arg := range args[1:] {
		if f(arg, res) == 1 {
			res = arg
		}
	}

	return res
}

// SearchOverlaps returns a slice of all intervals that overlaps with i
func (t *Tree[C, T]) SearchOverlaps(i Interval[C, T]) []Interval[C, T] {
	var res []Interval[C, T]
	searchOverlaps(t.root, i, &res)

	return res
}

func searchOverlaps[C Comparator[T], T any](n *node[C, T], i Interval[C, T], res *[]Interval[C, T]) {
	if n == nil {
		return
	}

	var c C
	cmpFunc := c.Compare

	if cmpFunc(i.a, n.max) == 1 {
		return
	}

	searchOverlaps(n.left, i, res)

	if overlaps(n.val, i) {
		*res = append(*res, n.val)
	}

	if cmpFunc(i.b, n.max) == -1 {
		return
	}

	searchOverlaps(n.right, i, res)
}

func overlaps[C Comparator[T], T any](i1, i2 Interval[C, T]) bool {
	res := i1.Relate(i2)
	return res != NoRel && res != Before && res != After
}
