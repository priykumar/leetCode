/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    q []int
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "nil"
    }

    q := []*TreeNode{root}
    res := ""
    for len(q)>0 {
        sz:=len(q)

        for i:=0; i<sz; i++ {
            if q[i] != nil {
                res = res + fmt.Sprintf("_%d", q[i].Val)
            } else {
                res = res + "_nil"
            }
            if q[i] != nil {
                q = append(q, q[i].Left)
                q = append(q, q[i].Right)
            }
        }
        q=q[sz:]
    }

    fmt.Println(res[1:])
    return res[1:]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {  
    getNumber := func(s string) (int) {
        if strings.Contains(s, "-") {
            s=s[1:]
            res, _ := strconv.Atoi(s)
            return res * -1
        } else {
            res, _ := strconv.Atoi(s)
            return res
        }
    } 

    if data == "nil" {
        return nil
    }

    a := strings.Split(data, "_")
    rootVal := getNumber(a[0])
    root := &TreeNode{rootVal, nil, nil}
    q := []*TreeNode{root}

    currA_Pos:=1
    for currA_Pos < len(a) {
        sz := len(q)
        for i:=0; i<sz && currA_Pos < len(a); i++ {
            leftnewNode := &TreeNode{-1, nil, nil}
            rightnewNode := &TreeNode{-1, nil, nil}

            fmt.Println("Number is", a[currA_Pos])
            if a[currA_Pos] != "nil" {
                num := getNumber(a[currA_Pos])
                leftnewNode.Val = num
            } else {
                leftnewNode = nil
            }

            fmt.Println("Number is", a[currA_Pos+1])
            if a[currA_Pos+1] != "nil" {
                num := getNumber(a[currA_Pos+1])
                rightnewNode.Val = num
            } else {
                rightnewNode = nil
            }

            if q[i] != nil {
                fmt.Println("left", q[i], leftnewNode)
                q[i].Left = leftnewNode
                q = append(q, leftnewNode)
                currA_Pos++

                fmt.Println("right", q[i], rightnewNode)
                q[i].Right = rightnewNode
                q = append(q, rightnewNode)
                currA_Pos++
            } 
            // else if q[i] != nil && i%2 == 1 {
            //     fmt.Println("right", q[i], newNode)
            //     q[i].Right = newNode
            //     q = append(q, newNode)
            //     currA_Pos++
            // }
        }
        q=q[sz:]
    }


    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */