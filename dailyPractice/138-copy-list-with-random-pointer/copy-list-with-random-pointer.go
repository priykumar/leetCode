/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
  // create random list
  if head == nil {
    return nil
  }
  curr := head
  for curr != nil {
    newNode := &Node{
      Val: curr.Val,
      Next: curr.Next,
      Random: curr.Random,
    }

    curr.Next = newNode
    curr = curr.Next.Next
  }

  // populate random
  randomHead := head.Next
  randomcurr := randomHead
  curr = head
  for curr != nil {
    if curr.Random != nil {
      randomcurr.Random = curr.Random.Next
    }
    curr = curr.Next.Next
    if randomcurr.Next != nil {
      randomcurr = randomcurr.Next.Next
    }
  }

  // // populate random
  // curr = head
  // for curr != nil {
  //   if curr.Random != nil {
  //     curr.Next.Random = curr.Random.Next
  //   }
  //   curr = curr.Next.Next
  // }

  // separate out random list
  randomHead = head.Next
  randomcurr = randomHead
  curr = head
  for curr != nil {
    curr.Next = curr.Next.Next
    if randomcurr.Next != nil {
      randomcurr.Next = randomcurr.Next.Next
    }

    curr = curr.Next
    randomcurr = randomcurr.Next
  }

  return randomHead
}