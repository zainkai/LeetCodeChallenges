/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode[]} lists
 * @return {ListNode}
 */
var mergeKLists = function(lists) {
  let resHead = new ListNode(null)
  let curr = resHead
  while(true) {
      const isAllNull = lists.reduce((acc,l) => acc = acc && (l == null), true)
      if(isAllNull) break
      let val = Infinity; let temp
      lists.forEach(head => {
          if(head && head.val < val) {
              temp = head
              val = temp.val
          }
      })
      const idx = lists.findIndex(h => h === temp)
      lists[idx] = lists[idx].next
      curr.next = temp
      curr = curr.next
  }
  return resHead.next
};