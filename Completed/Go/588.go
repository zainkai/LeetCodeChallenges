import (
    "strings"
    "sort"
)

type Node struct {
    name string
    isFile bool
    identifier int
    children map[string]*Node
}

type FileSystem struct {
    contentDb map[int]string
    root *Node
    identifierCount int
}


func Constructor() FileSystem {
    return FileSystem{
        contentDb: map[int]string{},
        root: &Node{
            name: "",
            isFile: false,
            identifier: 0,
            children: map[string]*Node{},
        },
        identifierCount: 0,
    }
}

// helper
func (this *FileSystem) travDir(path string, doCreate bool) *Node {
    pathArr := this.normalizePath(path)
    
    currNode := this.root
    for _, n := range pathArr {
        if nextNode, ok := currNode.children[n]; !ok && doCreate {
            currNode.children[n] = &Node{
                name: n,
                isFile: false,
                identifier: 0,
                children: map[string]*Node{},
            }
            currNode = currNode.children[n]
        } else {
            currNode = nextNode
        }
    }
    
    return currNode
}

// helper
func (this *FileSystem) normalizePath (path string) []string {
    pathArr := strings.Split(path, "/")
    stk := []string{}
    for _, p := range pathArr {
        if p == "." || p == "" {
            continue
        } else if p == ".." {
            if len(stk) > 0 {
                stk = stk[:len(stk)-1]
            }
        } else {
            stk = append(stk, p)
        }
    }
    
    return stk
}


func (this *FileSystem) Ls(path string) []string {
    currNode := this.travDir(path, false)
    if currNode.isFile {
        return []string{currNode.name}
    }
    
    res := []string{}
    for childNames := range currNode.children {
        res = append(res, childNames)
    }
    
    
    sort.Slice(res, func(i, j int) bool {
        return strings.Compare(res[i], res[j]) < 1
    })
    return res
}


func (this *FileSystem) Mkdir(path string)  {
    this.travDir(path, true)
}


func (this *FileSystem) AddContentToFile(filePath string, content string)  {
    currNode := this.travDir(filePath, true)
    if !currNode.isFile {
        currNode.isFile = true
        this.identifierCount++
        currNode.identifier = this.identifierCount
    }
    
    // add to content Db
    this.contentDb[currNode.identifier] += content
}


func (this *FileSystem) ReadContentFromFile(filePath string) string {
    currNode := this.travDir(filePath, false)
    idnt := currNode.identifier
    
    return this.contentDb[idnt]
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
