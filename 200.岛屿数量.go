/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	//
	if len(grid) == 0 {
		return 0
	}
	//
	var island int = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				island++
				dfsToZero(grid, i, j)
			}
		}
	}
	//
	return island
}

func dfsToZero(grid [][]byte, i, j int) {
	//
	grid[i][j] = '0'
	//
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfsToZero(grid, i-1, j)
	}
	//
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		dfsToZero(grid, i+1, j)
	}
	//
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfsToZero(grid, i, j-1)
	}
	//
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		dfsToZero(grid, i, j+1)
	}
}

// @lc code=end

