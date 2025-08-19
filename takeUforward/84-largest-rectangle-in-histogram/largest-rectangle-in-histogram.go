func largestRectangleArea(heights []int) int {
    n:=len(heights)
    stack:=[]int{}
    left, right := make([]int, n), make([]int, n)

    // find min at left
    for i:=0; i<n; i++ {
        for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
            stack = stack[:len(stack)-1]
        }

        if len(stack) == 0{
            left[i] = -1
        } else {
            left[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }

    stack = []int{}
    // find min at left
    for i:=n-1; i>=0; i-- {
        for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
            stack = stack[:len(stack)-1]
        }

        if len(stack) == 0{
            right[i] = n
        } else {
            right[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }

    fmt.Println(left, right)
    
    res := 0
    for i:=0; i<n; i++ {
        width := right[i]-left[i]-1
        res = max(res, width*heights[i])
    }

    return res
}