import (
    "strings"
)

/**
 * // This is HtmlParser's API interface.
 * // You should not implement it, or speculate about its implementation
 * type HtmlParser struct {
 *     func GetUrls(url string) []string {}
 * }
 */

func crawl(startUrl string, htmlParser HtmlParser) []string {
    targetHost := getHostName(startUrl)
    visited := map[string]bool{}
    res := []string{}
    
    q := []string{startUrl}
    for len(q) > 0 {
        top := q[0] // get first in queue
        q = q[1:] // remove first in queue
        
        // skip url
        if getHostName(top) != targetHost {
            continue
        } else if _, ok := visited[top]; ok {
            continue
        }
        
        // visit & add to res
        visited[top] = true
        res = append(res, top)
        
        
        // get next
        next := htmlParser.GetUrls(top)
        q = append(q, next...)
    }
    
    return res
}

func getHostName(url string) string {
    startIdx := strings.LastIndex(url, "//")+2
    endIdx := startIdx
    for ; endIdx < len(url); endIdx++ {
        if url[endIdx] == '/' || url[endIdx] == ':' {
            break
        }
    }
    
    return url[startIdx:endIdx]
}
