// SC: O(n)
// func majorityElement(nums []int) int {
//     n:=len(nums)
//     eleCount:=map[int]int{}
//     for _, ele := range nums {
//         eleCount[ele]++
//         if eleCount[ele] > n/2 {
//             return ele
//         }
//     }

//     return -1
// }

// SC: O(1)
// Boyer–Moore Voting Algorithm
func majorityElement(nums []int) int {
    count, res := 0, 0
    for _, ele := range nums {
        if count==0 {
            count++
            res=ele
        } else if ele == res {
            count++
        } else {
            count--
        }
    }
    return res
}