Children Sum Property in Binary Tree

Given the root of a binary tree, return true if and only if every node’s value is equal to the sum of the values stored in its left and right children.

For any missing ( null ) child, its value is treated as 0.
A leaf node automatically satisfies the rule because both children are null.

Examples:
Input: root = [1,4,3,5]

Output: false

Explanation:

The root is 1, but its children sum to 4 + 3 = 7. Since 1 ≠ 7, the tree violates the property.
Input: root = [10,4,6,1,3,2,4]

Output: true

Explanation:

4 = 1 + 3
6 = 2 + 4
10 = 4 + 6
All internal nodes satisfy the condition.