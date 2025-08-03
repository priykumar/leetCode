Given a root of binary search tree and a key(node) value, find the floor and ceil value for that particular key value.

Floor Value Node: Node with the greatest data lesser than or equal to the key value. 
Ceil Value Node: Node with the smallest data larger than or equal to the key value.
If a particular floor or ceil value is not present then output -1.


Input : root = [8, 4, 12, 2, 6, 10, 14] , key = 11

Output : [10, 12]


Input : root = [8, 4, 12, 2, 6, 10, 14] , key = 15

Output : [14, -1]
