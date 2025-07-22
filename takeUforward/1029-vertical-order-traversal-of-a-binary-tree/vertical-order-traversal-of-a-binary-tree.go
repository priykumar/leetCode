/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(A *TreeNode)  ([][]int) {
    distanceFromCenter_nodes := map[int][]int{}
    temp_distanceFromCenter_nodes := map[int][]int{}
    minDist, maxDist := 0, 0
    res := [][]int{}
    if A == nil {
        return res
    }

    type node_dist_pair struct {
        node *TreeNode
        dist int
    }
    queue := []node_dist_pair{}
    queue = append(queue, node_dist_pair{A, 0})
    distanceFromCenter_nodes[0] = []int{A.Val}
    
    for len(queue) > 0 {
        sz := len(queue)
        // fmt.Println(queue)
        temp_distanceFromCenter_nodes = map[int][]int{}
        for i:=0; i<sz; i++ {
            dist := queue[i].dist
            node := queue[i].node
            // distanceFromCenter_nodes[dist] = append(distanceFromCenter_nodes[dist], node.Val)

            minDist = min(minDist, dist)
            maxDist = max(maxDist, dist)

            if node.Left != nil {
                queue = append(queue, node_dist_pair{node.Left, dist-1})
                temp_distanceFromCenter_nodes[dist-1] = append(temp_distanceFromCenter_nodes[dist-1], node.Left.Val)
            }
            if node.Right != nil {
                queue = append(queue, node_dist_pair{node.Right, dist+1})
                temp_distanceFromCenter_nodes[dist+1] = append(temp_distanceFromCenter_nodes[dist+1], node.Right.Val)
            }
        }
        // fmt.Println("here", queue)
        for key, value := range temp_distanceFromCenter_nodes {
            sort.Ints(value)
            if _, exist := distanceFromCenter_nodes[key]; exist {
                distanceFromCenter_nodes[key] = append(distanceFromCenter_nodes[key], value...)
            } else {
                distanceFromCenter_nodes[key] = value
            }
        }

        fmt.Println(distanceFromCenter_nodes)
        queue = queue[sz:]
    }

    for i:=minDist; i<=maxDist; i++ {
        // sort.Ints(distanceFromCenter_nodes[i])
        res = append(res, distanceFromCenter_nodes[i])
    }
    //fmt.Println(distanceFromCenter_nodes)
    // for _, value := range distanceFromCenter_nodes {
    //     res = append(res, value[0])
    // }

    return res
}
