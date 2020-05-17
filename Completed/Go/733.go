type Point struct {
    x,y int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    oldColor := image[sr][sc]
    visited := map[Point]bool{}
    
    dfs(image, sr,sc, oldColor, newColor, visited)
    
    return image
}

func dfs(image [][]int, x,y, oldColor, newColor int, visited map[Point] bool) {
    if x < 0 || y < 0 || x > len(image)-1 || y > len(image[x])-1 {
        return
    } else if visited[Point{x,y}] {
      return  
    } else if image[x][y] != oldColor {
        return
    } else {
        image[x][y] = newColor
        visited[Point{x,y}]= true
    }
    
    dfs(image, x-1, y, oldColor, newColor, visited)
    dfs(image, x+1, y, oldColor, newColor,visited)
    dfs(image, x, y-1, oldColor, newColor,visited)
    dfs(image, x, y+1, oldColor, newColor,visited)
}