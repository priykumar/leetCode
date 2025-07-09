func merge(nums1 []int, m int, nums2 []int, n int)  {
    i1, i2 := m-1, 0
    for i1 >= 0 && i2 < n && nums1[i1]>=nums2[i2] {
        nums1[i1], nums2[i2] = nums2[i2], nums1[i1]
        i1--
        i2++
    }

    sort.Ints(nums1[:m])
    sort.Ints(nums2)

    for i:=0; i<len(nums2); i++ {
        nums1[m+i] = nums2[i]
    }

}