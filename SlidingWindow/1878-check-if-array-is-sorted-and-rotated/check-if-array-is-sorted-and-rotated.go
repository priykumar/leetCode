func check(nums []int) bool {
    n, i := len(nums), 1
    if n==1 || n==2 {
        return true
    }
    for i < n && nums[i-1] <= nums[i] {
        i++
    }

    if i == n {
        return true
    }

    j:=i+1
    for j<n && nums[j-1] <= nums[j] {
        j++
    } 
    if j==n && nums[n-1] <= nums[0] {
        return true
    } 

    return false 


}