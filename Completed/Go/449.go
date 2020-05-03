import (
	"strconv"
	"strings"
)


type Codec struct{}

func Constructor() Codec { return Codec{} }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := []string{}
	serializeHelper(root, &result)

	return strings.Join(result, ",")
}
func serializeHelper(root *TreeNode, arr *[]string) {
	if root == nil {
		return
	}
	*arr = append(*arr, strconv.Itoa(root.Val))
	serializeHelper(root.Left, arr)
	serializeHelper(root.Right, arr)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	dataArr := strings.Split(data, ",")
	i := 0
    
	return deserializeHelper(dataArr, &i, -1<<31, 1<<32-1)
}
func deserializeHelper(data []string, i *int, left, right int) *TreeNode {
	if *i >= len(data) {
		return nil
	}

	val, err := strconv.Atoi(data[*i]) // note check for err
	if err != nil || val < left || val > right {
		return nil
	}

	node := &TreeNode{val, nil, nil}
	*i++
	node.Left = deserializeHelper(data, i, left, val)
	node.Right = deserializeHelper(data, i, val, right)

	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */