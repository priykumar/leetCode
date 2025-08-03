Given two integer arrays nums1 and nums2 and an integer k, return the maximum k valid sum combinations from all possible sum combinations using the elements of nums1 and nums2.



A valid sum combination is made by adding one element from nums1 and one element from nums2. Return the answer in non-increasing order.


Examples:
Input: nums1 = [7, 3], nums2 = [1, 6], k = 2

Output: [13, 9]

Explanation: The 2 maximum combinations are made by:

nums1[0] + nums2[1] = 13

nums1[1] + nums2[1] = 9

Input: nums1 = [3, 4, 5], nums2 = [2, 6, 3], k = 2

Output: [11, 10]

Explanation: The 2 maximum combinations are made by:

nums1[2] + nums2[1] = 11

nums1[1] + nums2[1] = 10