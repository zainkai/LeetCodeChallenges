// // first iteration
// var printLinkedListInReverse = function(head) {
//   arr = []
//   helper(head, arr)

//   for (let i = arr.length -1; i >= 0; i--) {
//     arr[i].printValue()
//   }
// };

// function helper(head, arr) {
// if (head == null) {
//   return
// }

// arr.push(head)

// helper(head.getNext(), arr)
// }

/**
 * // This is the ImmutableListNode's API interface.
 * // You should not implement it, or speculate about its implementation.
 * function ImmutableListNode() {
 *    @ return {void}
 *    this.printValue = function() { // print the value of this node.
 *       ...
 *    }; 
 *
 *    @return {ImmutableListNode}
 *    this.getNext = function() { // return the next node.
 *       ...
 *    };
 * };
 */

/**
 * @param {ImmutableListNode} head
 * @return {void}
 */
var printLinkedListInReverse = function(head) {
  arr = []
  for (let h = head; h != null; h = h.getNext()) {
    arr.push(h)
  }

  while (arr.length > 0) {
    arr.pop().printValue()
  }
};
