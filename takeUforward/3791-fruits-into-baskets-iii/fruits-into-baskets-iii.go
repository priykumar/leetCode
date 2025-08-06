// Complexity:
// Time O(NLogN) and Space O(N) where N is the length of
// fruits and baskets.
func numOfUnplacedFruits(fruits []int, baskets []int) int {
	tree := newMaxSegmentTree(baskets)
	result := 0
	for _, fruit := range fruits {
		if !tree.place(fruit) {
			result++
		}
	}
	return result
}

type maxSegmentTree struct {
	// The start indices for each level, from leaves to top.
	starts []int
    // The value of nodes, from leaves to root and level by
    // level.
	values []int
}

func newMaxSegmentTree(baskets []int) *maxSegmentTree {
	// Determine the start indices for each level
	start := 0
	levelStarts := make([]int, 0)
	for width := len(baskets); width > 1; width = (width + 1) / 2 {
		levelStarts = append(levelStarts, start)
		start += width
	}
	levelStarts = append(levelStarts, start)

	// Fill values from the leaves to the root
	values := make([]int, start+1)

	copy(values, baskets)
	for i := 1; i < len(levelStarts); i++ {
		prevStart := levelStarts[i-1]
		currStart := levelStarts[i]

		k := currStart
		for j := prevStart; j+1 < currStart; j += 2 {
			values[k] = max(values[j], values[j+1])
			k++
		}
		if (currStart-prevStart)&1 == 1 {
			values[k] = values[currStart-1]
		}
	}

	return &maxSegmentTree{levelStarts, values}
}

// place places the fruit in the first large enough basket.
// It returns false if such a basket does not exist.
func (tree *maxSegmentTree) place(fruit int) bool {
	path, ok := tree.query(fruit)
	if ok {
		tree.update(path, 0)
	}
	return ok
}

// query returns the path (indices) from the first large
// enough basket to the root. The boolean indicates whether
// such a basket exists.
func (tree *maxSegmentTree) query(fruit int) ([]int, bool) {
	index := len(tree.values) - 1
	if tree.values[index] < fruit {
		return nil, false
	}

	// Walk from the root to the leaf
	path := make([]int, len(tree.starts))
	for i := len(tree.starts) - 1; i > 0; i-- {
		path[i] = index
		currStart := tree.starts[i]
		nextStart := tree.starts[i-1]

		index = nextStart + (index-currStart)*2
		if tree.values[index] < fruit {
			index++
		}
	}
	path[0] = index
	return path, true
}

// update updates the leafValue given a path from the root
// to that leaf.
func (tree maxSegmentTree) update(path []int, leafValue int) {
	// Walk from the leaf to the root
	value := leafValue
	for i := 0; i < len(tree.starts)-1; i++ {
		index := path[i]
		tree.values[index] = value

		sibIndex := index
		if (index & 1) == (tree.starts[i] & 1) {
			sibIndex++
		} else {
			sibIndex--
		}

		if sibIndex < tree.starts[i+1] {
			value = max(value, tree.values[sibIndex])
		}
		if value == tree.values[path[i+1]] {
			return
		}
	}
	tree.values[len(tree.values)-1] = value
}