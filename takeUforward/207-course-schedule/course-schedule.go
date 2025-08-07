// Find cycle in directed graph
func createAdjList(numCourses int, p [][]int) [][]int {
    adj := make([][]int, numCourses) 
    for i:=0; i<len(p); i++ {
        adj[p[i][0]] = append(adj[p[i][0]], p[i][1]) 
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

func canFinish(numCourses int, prerequisites [][]int) bool {
    adj := createAdjList(numCourses, prerequisites)

    vstd := make([]bool, numCourses)
    path := make([]bool, numCourses)
    for u:=0; u<numCourses; u++ {
        if !vstd[u] {
            if dfs(u, vstd, path,adj) {
                return false
            }
        }
    }

    return true
}