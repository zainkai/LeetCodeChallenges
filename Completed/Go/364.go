// Brute Force
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
//  func depthSumInverse(nestedList []*NestedInteger) int {
//     sum, depth := 0, 1

//     maxDepth(nestedList, 1, &depth)
//     dfs(nestedList, 1, depth, &sum)

//     return sum
// }

// func maxDepth(nl []*NestedInteger, depth int, mDepth *int) {
//     *mDepth = max(*mDepth, depth)
//     for _, l := range nl {
//         if !l.IsInteger() {
//             maxDepth(l.GetList(), depth+1, mDepth)
//         }
//     }
// }

// func dfs(nl []*NestedInteger, depth int, maxDepth int, sum *int) {
//     levelSum := 0

//     for _, list := range nl {
//         if list.IsInteger() {
//             levelSum += list.GetInteger()
//         } else {
//             dfs(list.GetList(), depth+1, maxDepth, sum)
//         }
//     }

//     *sum += levelSum * (maxDepth - depth+1)
// }

// func max (a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }