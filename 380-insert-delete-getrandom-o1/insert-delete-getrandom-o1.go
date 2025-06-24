type RandomizedSet struct {
	valToIndex map[int]int
	values     []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		valToIndex: make(map[int]int),
		values:     []int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.valToIndex[val]; exists {
		return false
	}
	rs.values = append(rs.values, val)
	rs.valToIndex[val] = len(rs.values) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	index, exists := rs.valToIndex[val]
	if !exists {
		return false
	}

	last := rs.values[len(rs.values)-1]
	rs.values[index] = last
	rs.valToIndex[last] = index

	rs.values = rs.values[:len(rs.values)-1]
	delete(rs.valToIndex, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(rs.values))
	return rs.values[randomIndex]
}