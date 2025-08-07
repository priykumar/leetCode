/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
// a := []int{1,2,3}
// b := a // shallow copy
// a[1] = 4
// b: [1,4,3]

// time: O(n)
func cloneGraph(node *Node) *Node {
    if node == nil{
        return node
    }
    queue := []*Node{node}
    nodes := []*Node{}
    copyNodes := map[*Node]*Node{}
    visited := map[*Node]bool{}
    
    for len(queue) > 0{
        queueLen := len(queue)
        
        for i := 0; i < queueLen; i++{
            cur := queue[i]
            if visited[cur]{
                continue
            }
            
            visited[cur] = true
            nodes = append(nodes, cur)
            copyNodes[cur] = &Node{
                Val: cur.Val,
            }
            //neighbor nodes
            for _, nbr := range cur.Neighbors{
                queue = append(queue, nbr)
            }
        }
        
        queue = queue[queueLen:]
    }

    for _, node := range nodes{
        copyNodes[node].Neighbors = []*Node{}
        for _, nbr := range node.Neighbors{
            copyNodes[node].Neighbors = append(copyNodes[node].Neighbors, copyNodes[nbr]) 
        }
    }
    
    return copyNodes[node]
}