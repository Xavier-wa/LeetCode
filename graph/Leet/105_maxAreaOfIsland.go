package leet

func maxAreaOfIsland(grid [][]int) int {
	var res int = 0
	vis := make([][]bool, len(grid))
	for i, _ := range vis {
		vis[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if vis[i][j] == false && grid[i][j] == 1 {
				cutAres := 0
				vis[i][j] = true
				cutAres ++
				dfs_ares(grid, vis, i, j, &cutAres)
				res = max(res, cutAres)
			}
		}
	}
	return res 
}
func dfs_ares(grid [][]int, vis [][]bool, x, y int, cutAres *int) {
	var dir [][]int = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	for i := 0; i < 4; i++ {
		nextX := x + dir[i][0]
		nextY := y + dir[i][1]
		if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
			continue
		}
		if grid[nextX][nextY] == 1 && vis[nextX][nextY] == false {
			vis[nextX][nextY] = true
			*cutAres ++
			dfs_ares(grid, vis, nextX, nextY, cutAres)
		}
	}
}
