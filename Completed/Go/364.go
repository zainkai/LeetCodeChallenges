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
func depthSumInverse(nestedList []*NestedInteger) int {
    maxDepth := 0
    depthMap := map[int][]int{}
    
    q := nestedList
    depth := 1
    for len(q) > 0 {
        end := len(q)
        for i := 0; i < end; i++ {
            top := q[0]
            q = q[1:]
            
            if top.IsInteger() {
                x := top.GetInteger()
                maxDepth = max(maxDepth, depth)
                depthMap[x] = append(depthMap[x], depth)
            } else {
                q = append(q, top.GetList()...)
            }
        }
        depth++
    }
    
    res := 0
    for key, arr := range depthMap {
        for _, depth := range arr {
            weight := maxDepth - depth + 1
            res += key * weight
        }
    }
    
    return res
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
