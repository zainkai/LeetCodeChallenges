function getIntersectionNode(headA, headB) {
  let iterA = headA, iterB = headB
  
  while(iterA !== iterB) {
    if (iterA === null) {
      iterA = headB
    } else { iterA = iterA.next }
    if (iterB === null) {
      iterB = headA
    } else { iterB = iterB.next }
  }
  return iterA
}

function getIntersectionNode(headA, headB, iterA=headA,iterB=headB) {
  if(iterA === iterB) {
    return iterA
  } else if(iterA === null) {
    return getIntersectionNode(headA,headB, headB, iterB.next)
  } else if(iterB === null) {
    return getIntersectionNode(headA,headB, iterA.next, headA)
  } else {
    return getIntersectionNode(headA,headB, iterA.next, iterB.next)
  }
}
