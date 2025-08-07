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

func isBipartite(graph [][]int) bool {
    nodeCount := len(graph)
    color := make([]int, nodeCount)
    for i:=0; i<nodeCount; i++ {
        color[i]=-1
    }

    for i:=0; i<nodeCount; i++ {
        if color[i]==-1 {
            if !bfs(i, color, graph) {
                return false
            }
        }
    }

    return true
}