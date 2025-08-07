// Find cycle in directed graph
func createAdjList(numCourses int, p [][]int) [][]int {
    adj := make([][]int, numCourses) 
    for i:=0; i<len(p); i++ {
        adj[p[i][1]] = append(adj[p[i][1]], p[i][0]) 
    }

    return adj
}

func dfs(u int, vstd, path []bool, adj [][]int) bool {
    vstd[u] = true
    path[u] = true
    for _, v := range adj[u] {
        if !vstd[v] {
            if dfs(v, vstd, path, adj) {
                return true
            }
        } else if path[v] {
            return true // cycle detected
        }
    }

    path[u] = false
    return false
}

func DFS_Approach(numCourses int, adj [][]int) bool{
    vstd := make([]bool, numCourses)
    path := make([]bool, numCourses)
    for u:=0; u<numCourses; u++ {
        if !vstd[u] {
            if dfs(u, vstd, path, adj) {
                return false
            }
        }
    }

    return true
}

func BFS_Approach(numCourses int, adj [][]int) bool {
    indegree := make([]int, numCourses)
    for i:=0; i<len(adj); i++ {
        for j:=0; j<len(adj[i]); j++ {
            indegree[adj[i][j]]++
        }
    }

    fmt.Println(adj, indegree)
    q := []int{}
    for i, ele := range indegree {
        if ele == 0{
            q=append(q, i)
        }
    }

    count := 0
    tsort := []int{}
    for len(q) > 0 {
        u := q[0]
        fmt.Println(u)
        q = q[1:]
        count++
        tsort = append(tsort, u)
        for _, v := range adj[u] {
            indegree[v]--
            if indegree[v] == 0 {
                q = append(q, v)
            }
        }
    }

    fmt.Println(count, tsort)
    return count == numCourses
}

func canFinish(numCourses int, prerequisites [][]int) bool {
    adj := createAdjList(numCourses, prerequisites)

    return BFS_Approach(numCourses, adj)
    // return DFS_Approach(numCourses, adj)
}