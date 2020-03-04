// 在给定的网格中，每个单元格可以有以下三个值之一：
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
// 每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
// 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。

// 示例 1：
// 输入：[[2,1,1],[1,1,0],[0,1,1]]
// 输出：4

// 示例 2：
// 输入：[[2,1,1],[0,1,1],[1,0,1]]
// 输出：-1
// 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。

// 示例 3：
// 输入：[[0,2]]
// 输出：0
// 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
//
// 提示：
// 1 <= grid.length <= 10
// 1 <= grid[0].length <= 10
// grid[i][j] 仅为 0、1 或 2
package easy

func OrangesRotting(grid [][]int) int {
	// 1. 找到所有腐烂的橘子
	fresh := find(grid, 1)
	if len(fresh) == 0 {
		return 0
	}
	rotted := find(grid, 2)
	if len(rotted) == 0 {
		return -1
	}
	// 2. 开始腐烂
	minutes := 0
	canRot := findCanRot(grid)
	for len(canRot) != 0 {
		for k := range canRot {
			grid[canRot[k][0]][canRot[k][1]] = 2
		}
		minutes++
		canRot = findCanRot(grid)
	}
	if len(find(grid, 1)) != 0 {
		return -1
	}
	return minutes
}
func find(grid [][]int, target int) [][]int {
	var result [][]int
	for k := range grid {
		for j := range grid[0] {
			if grid[k][j] == target {
				tmp := []int{k, j}
				result = append(result, tmp)
			}
		}
	}
	return result
}
func findCanRot(grid [][]int) [][]int {
	var result [][]int
	for k := range grid {
		for j := range grid[0] {
			if grid[k][j] == 1 {
				flag := false
				// 上
				if k-1 >= 0 && grid[k-1][j] == 2 {
					flag = true
				}
				// 下
				if k+1 <= len(grid)-1 && grid[k+1][j] == 2 {
					flag = true
				}
				// 左
				if j-1 >= 0 && grid[k][j-1] == 2 {
					flag = true
				}
				// 右
				if j+1 <= len(grid[0])-1 && grid[k][j+1] == 2 {
					flag = true
				}
				if flag {
					tmp := []int{k, j}
					result = append(result, tmp)
				}
			}
		}
	}
	return result
}
