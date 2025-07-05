package leet

var dir [][]int = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func numIslands(grid [][]byte) int {
	var res int = 0
	vis := make([][]bool, len(grid))
	for i, _ := range vis {
		vis[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if vis[i][j] == false && grid[i][j] == '1' {
				dfs(grid, vis, i, j, res)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, vis [][]bool, x, y, res int) {
	vis[x][y] = true
	for i := 0; i < 4; i++ {
		nextX := x + dir[i][0]
		nextY := y + dir[i][1]
		if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
			continue
		}
		if grid[nextX][nextY] == '1' && vis[nextX][nextY] == false {
			dfs(grid, vis, nextX, nextY, res)
		}
	}
}

//bfs

type arix struct {
	x int
	y int
}

func numIslands2(grid [][]byte) int {

	res := 0
	vis := make([][]bool, len(grid))
	for i, _ := range vis {
		vis[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !vis[i][j] && grid[i][j] == '1' {
				bfs(vis, grid, i, j)
				res++
			}
		}
	}

	return res
}

func bfs(vis [][]bool, grid [][]byte, x, y int) {
	if vis[x][y] {
		return
	}
	que := []arix{}

	que = append(que, arix{x, y})

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		for i := 0; i < 4; i++ {
			nextX := cur.x + dir[i][0]
			nextY := cur.y + dir[i][1]
			if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
				continue
			}
			if grid[nextX][nextY] == '1' && vis[nextX][nextY] == false {
				que = append(que, arix{nextX, nextY})
				vis[nextX][nextY] = true
			}
		}
	}
}
