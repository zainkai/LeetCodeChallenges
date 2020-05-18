type node struct {
    isWord bool
    conns map[byte]*node
}

type WordDictionary struct {
    trie *node
}

func newNode(isWord bool) *node {
    return &node{
        isWord,
        map[byte]*node{},
    }
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{
        trie: newNode(false),
    }
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
    curr := this.trie
    
    for _, r := range word {
        b := byte(r)
        
        if _, ok := curr.conns[b]; !ok {
            curr.conns[b] = newNode(false)
        }
        curr = curr.conns[b]
    }
    
    curr.isWord = true
}

type frame struct {
    n *node
    idx int
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    q := []frame{
        frame{this.trie, 0},
    }
    
    for len(q) > 0 {
        top := q[0]
        q = q[1:]
        
        if top.n.isWord && top.idx == len(word) {
            return true
        } else if top.idx >= len(word) {
            continue
        } else if word[top.idx] != '.' {
            char := word[top.idx]
            if next, ok := top.n.conns[char]; ok {
                q = append(q, frame{
                    next,
                    top.idx+1,
                })
            }
        } else {
            for _, conn := range top.n.conns {
                q = append(q, frame{
                    conn,
                    top.idx+1,
                })
            }   
        }
    }
    return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */