type NodeAndIndex struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := []NodeAndIndex{{node: root, index: 0}}

	var maxWidth int

	for len(queue) != 0 {
		levelSize := len(queue)
		levelIndex := queue[0].index

		for i := 0; i < levelSize; i++ {
			// pop
			data := queue[0]
			queue = append(queue[:0], queue[1:]...)

			if data.node.Left != nil {
				queue = append(queue, NodeAndIndex{node: data.node.Left, index: data.index * 2})
			}
			if data.node.Right != nil {
				queue = append(queue, NodeAndIndex{node: data.node.Right, index: data.index*2 + 1})
			}

			maxWidth = max(maxWidth, data.index-levelIndex+1)
		}
	}

	return maxWidth
}

func max(nums ...int) int {
	maxNum := nums[0]
	for _, num := range nums[1:] {
		if maxNum < num {
			maxNum = num
		}
	}

	return maxNum
}