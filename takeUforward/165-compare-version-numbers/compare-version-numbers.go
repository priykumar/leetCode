func compareVersion(version1 string, version2 string) int {
    v1:=strings.Split(version1, ".")
    v2:=strings.Split(version2, ".")

    m, n := len(v1), len(v2)
    i:=0
    for i<min(m, n){
        v11, _ := strconv.Atoi(v1[i])
        v22, _ := strconv.Atoi(v2[i])
        if v11 > v22 {
            return 1
        } else if v11 < v22 {
            return -1
        }

        i++
    }

    for i<m {
        if val, _ := strconv.Atoi(v1[i]); val != 0 {
            return 1
        }
        i++
    }

    for i<n {
        if val, _ := strconv.Atoi(v2[i]); val != 0 {
            return -1
        }
        i++
    }

    return 0
}