// brute force O(N log N)
// better O(NK) compare one by one
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 import (
    "sort"
)

func mergeKLists(lists []*ListNode) *ListNode {
    arr := []*ListNode{}
    for _, list := range lists {
        for node := list; node != nil; node = node.Next {
            arr = append(arr, node)
        }
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].Val < arr[j].Val
    })
    
    for i, node := range arr {
        if i+1 == len(arr) {
            node.Next = nil
        } else {
            node.Next = arr[i+1]
        }
    }
    
    if len(arr) == 0 {
        return nil
    }
    return arr[0]
}

