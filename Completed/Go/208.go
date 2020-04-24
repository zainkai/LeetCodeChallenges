type Trie struct {
    root *node
}

type node struct {
    isWord bool
    children map[rune]*node
}

func newNode() *node {
    return &node{
        false,
        make(map[rune]*node),
    }
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        newNode(),
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    curr := this.root
    for _, r := range word {
        if _, ok := curr.children[r]; !ok {
            curr.children[r] = newNode()
        }
        curr = curr.children[r]
    }
    curr.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    curr := this.root
    for _, r := range word {
        if _, ok := curr.children[r]; !ok {
            return false
        }
        curr = curr.children[r]
    }
    return curr.isWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    curr := this.root
    for _, r := range prefix {
        if _, ok := curr.children[r]; !ok {
            return false
        }
        curr = curr.children[r]
    }
    return curr.isWord || len(curr.children) != 0
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */