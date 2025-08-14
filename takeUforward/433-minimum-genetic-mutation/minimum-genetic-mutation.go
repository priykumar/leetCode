func isMutate(a, b string) bool {
    count:=0
    for i:=0; i<len(a); i++ {
        if a[i] != b[i] {
            count++
        }
        if count>1 {
            return false
        }
    }

    return count==1
}
func minMutation(startGene string, endGene string, bank []string) int {
    if startGene == endGene {
        return 0
    }

    bank = append([]string{startGene}, bank...)
    adj := make([][]int, len(bank))

    src, dst := 0, -1
    // Create ADJ matrix
    for i:=0; i<len(bank); i++ {
        // adj[i] = []int{}
        if bank[i] == endGene {
            dst = i
        }
        for j:=i; j<len(bank); j++ {
            if isMutate(bank[i], bank[j]) {
                adj[i] = append(adj[i], j)
                adj[j] = append(adj[j], i)
            }
        }
    }
    if dst == -1 {
        return -1
    }

    fmt.Println(src, dst, adj)

    // Find shortest path b/w src and dst
    type S struct {
        node int
        dist int
    }

    queue := []S{{src,0}}
    // vstd := make([]bool, len(bank))
    distance := make([]int, len(bank))
    for i:=0; i<len(bank); i++ {
        distance[i] = 1<<32
    }

    // vstd[0] = true
    for len(queue) > 0{
        unode, udist := queue[0].node, queue[0].dist
        queue=queue[1:]
    
        for _, v := range adj[unode] {
            if udist+1 < distance[v] {
                distance[v] = udist+1
                // vstd[v]=true
                queue=append(queue, S{v, udist+1})
            }
        }
    }

    if distance[dst] == 1<<32 {
        return -1
    }

    return distance[dst]
    
    return 0
}