func bfs(node int, color []int, graph [][]int) bool{
    color[node] = 1
    q := []int{node}
    for len(q) > 0 {
        u:=q[0]
        q=q[1:]

        for _, v := range graph[u] {
            if color[v] == -1 {
                color[v]=1-color[u]
                q=append(q, v)
            } else if color[v] == color[u] {
                return false
            }
        }
    }

    return true
}

func isBipartite_BFS(nodeCount int, color []int, graph [][]int) bool {
    for i:=0; i<nodeCount; i++ {
        if color[i]==-1 {
            if !bfs(i, color, graph) {
                return false
            }
        }
    }

    return true
}

func dfs(u int, color []int, graph [][]int) bool {
    for _, v := range graph[u] {
        if color[v] == -1 {
            color[v] = 1-color[u]
            if !dfs(v, color, graph) {
                return false
            }
        } else if color[v] == color[u] {
                return false
        }
    }

    return true
}

func isBipartite_DFS(nodeCount int, color []int, graph [][]int) bool {
    for i:=0; i<nodeCount; i++ {
        if color[i]==-1 {
            color[i]=1
            if !dfs(i, color, graph) {
                return false
            }
        }
    }

    return true
}

func isBipartite(graph [][]int) bool {
    nodeCount := len(graph)
    color := make([]int, nodeCount)
    for i:=0; i<nodeCount; i++ {
        color[i]=-1
    }

    return isBipartite_DFS(nodeCount, color, graph)
}