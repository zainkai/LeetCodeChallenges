class ListNode {
  constructor (next) {
    this.val = val
    this.next = next
  }
}

function addLLNums (l1, l2) { 
  let rootRes = new ListNode(null)
  let carry = 0
  let currl1 = l1, currl2 = l2, currRes = rootRes
  
  while (currl1 !== null && currl2 !== null) {
     let value = (currl1.val || 0) + (currl2.val || 0) + carry
     if(value >= 10) {
       carry = 1
       value = value % 10
     } else {
       carry = 0
     }

     currRes.val = value
     currRes.next = new ListNode(null)
     currRes = currRes.next

     currl1 = currl1.next
     currl2 = currl2.next
  }
  return rootRes
} 
