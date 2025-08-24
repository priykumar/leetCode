func createAdjList(numCourses int, prerequisites [][]int) [][]int{
    adj := make([][]int, numCourses)
    for i:=0; i<len(prerequisites); i++ {
        adj[prerequisites[i][1]] = append(adj[prerequisites[i][1]], prerequisites[i][0])
    }
    return adj
}
func findOrder(numCourses int, prerequisites [][]int) []int {
    adj := createAdjList(numCourses, prerequisites)
    indegree := make([]int, numCourses)
    for i:=0; i<len(adj); i++ {
        for _, v := range adj[i] {
            indegree[v]++
        }
    }

    queue := []int{}
    for i:=0; i<len(indegree); i++ {
        if indegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    res := make([]int, 0, numCourses)
    count:=0
    for len(queue)>0 {
        u := queue[0]
        queue = queue[1:]
        res = append(res, u)
        count++
        for _, v := range adj[u] {
            indegree[v]--
            if indegree[v] == 0 {
                queue = append(queue, v)
            }
        }
    }

    if count == numCourses {
        return res
    }
    return []int{}
}