/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    arr := []*ListNode{}
    
    curr := head
    for curr != nil {
        arr = append(arr, curr)
        curr = curr.Next
    }
    
    if len(arr) <= 1 {
        return
    }
    
    i, j := 0, len(arr)-1
    for i < j {
        arr[i].Next = arr[j]
        i++
        arr[j].Next = arr[i]
        j--
    }
    arr[i].Next = nil
}
