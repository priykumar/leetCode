type Class struct {
    pass, total int
    gain        float64
}

type MaxHeap []Class

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool  { return h[i].gain > h[j].gain } // max-heap
func (h *MaxHeap) Push(x any)         { *h = append(*h, x.(Class)) }
func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    ele := old[n-1]
    *h = old[:n-1]
    return ele
}


func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    h := &MaxHeap{}
    heap.Init(h)

    // Initialize heap with all classes
    for _, c := range classes {
        p, t := c[0], c[1]
        g := gain(p, t)
        heap.Push(h, Class{pass: p, total: t, gain: g})
    }

    // Assign extra students greedily
    for extraStudents > 0 {
        c := heap.Pop(h).(Class)
        c.pass++
        c.total++
        c.gain = gain(c.pass, c.total)
        heap.Push(h, c)
        extraStudents--
    }

    // Compute final average
    sum := 0.0
    for h.Len() > 0 {
        c := heap.Pop(h).(Class)
        sum += float64(c.pass) / float64(c.total)
    }

    return sum / float64(len(classes))
}

func gain(p, t int) float64 {
    return float64(p+1)/float64(t+1) - float64(p)/float64(t)
}
