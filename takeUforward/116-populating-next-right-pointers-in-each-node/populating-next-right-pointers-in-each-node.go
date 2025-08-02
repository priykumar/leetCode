/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return root
    }

    q := []*Node{root}
    for len(q) > 0 {
        fmt.Println(q)
        sz := len(q)
        for i:=0; i<sz; i++ {
            fmt.Println(i, sz)
            if i != sz-1 {
                q[i].Next = q[i+1]
            } 

            if q[i].Left != nil {
                q=append(q, q[i].Left)
            }
            if q[i].Right != nil {
                q=append(q, q[i].Right)
            }
        }
        q=q[sz:]
    }

    return root
}