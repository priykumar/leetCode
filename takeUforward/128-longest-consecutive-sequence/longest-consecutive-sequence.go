// func longestConsecutive(nums []int) int {
//     m:=map[int]int{}

//     for _, v := range nums{
//         m[v]++
//     }

//     maxlen, count := 0, 0
//     for i:=0; i<len(nums); i++ {
//         if m[nums[i]-1] == 0 {
//             count=0
//             ele:=nums[i]
//             for m[ele]>0 {
//                 count++
//                 ele++
//             }

//             if count > maxlen {
//                 maxlen=count
//             }
// //             // maxlen=max(maxlen, count)
// //         }
// //     }

// //     return maxlen
// // }

// func longestConsecutive(A []int) int {
//     m:=map[int]int{}
//     for _, v:=range A{
//         m[v]++
//     }
    
//     result:=0
//     for i:=0; i<len(A); i++ {
//         if m[A[i]-1] == 0 {
//             tempLen:=0
//             startEle:=A[i]
//             for m[startEle]>0 {
//                 tempLen++ //=m[startEle]
//                 startEle++
//             }
//             if tempLen > result {
//                 result=tempLen
//             }
//         }
//     }
    
//     return result
// }

func longestConsecutive(nums []int) int {
    // Construct a set out of the nums array.
    numsSet := make(map[int]struct{})
    for _, n := range nums {
        numsSet[n] = struct{}{}
    }

    // The answer is stored here.
    maxSequenceLen := 0

    // Iterate through the set.
    for n := range numsSet {
        // We check if n-1 is in the set. If it is, then n is not the beginning of a sequence
        // and we go to the next number immediately.
        if _, ok := numsSet[n-1]; !ok {
            // Otherwise, we increment n in a loop to see if the next consecutive value is stored in nums.
            seqLen := 1
            for {
                if _, ok = numsSet[n+seqLen]; ok {
                    seqLen++
                    continue
                }
                // When the sequence is over, see if we did better than before.
                maxSequenceLen = max(seqLen, maxSequenceLen)
                break
            }
        }
    }

    return maxSequenceLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}