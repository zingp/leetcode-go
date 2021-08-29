package main

/*
最短路径和
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/

import (
	"fmt"
)

func Min(x, y int) int{
	if x<y {
		return x
	}
	return y
}

func minPathSumHelp(grid [][]int, r int, c int ) int {
	if r == 0 && c == 0 {
		return grid[0][0]
	} else if r==0 && c > 0{
		return minPathSumHelp(grid, r, c-1)+grid[r][c]
	} else if r>0 && c==0 {
		return minPathSumHelp(grid, r-1, c)+grid[r][c]
	} else {
		return Min(minPathSumHelp(grid, r-1, c)+grid[r][c], minPathSumHelp(grid, r, c-1)+grid[r][c])
	}
}
func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}
	r := len(grid)-1
	c := len(grid[0])-1
	return minPathSumHelp(grid, r, c)
}

// 递归很慢，改成遍历O(n^2)
func minPathSum2(grid [][]int) int {
	if grid == nil {
		return 0
	}
	rLen := len(grid)
	cLen := len(grid[0])
	for i:=0;i<rLen;i++ {
		for j:=0; j<cLen;j++ {
			if i>0 && j==0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			}else if i==0 && j>0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}else if i>0 && j>0 {
				grid[i][j] = grid[i][j] + Min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[rLen-1][cLen-1]
}

func main() {
	grid := [][]int {
	  {7,1,3,5,8,9,9,2,1,9,0,8,3,1,6,6,9,5},
      {9,5,9,4,0,4,8,8,9,5,7,3,6,6,6,9,1,6},
      {8,2,9,1,3,1,9,7,2,5,3,1,2,4,8,2,8,8},
      {6,7,9,8,4,8,3,0,4,0,9,6,6,0,0,5,1,4},
      {7,1,3,1,8,8,3,1,2,1,5,0,2,1,9,1,1,4},
      {9,5,4,3,5,6,1,3,6,4,9,7,0,8,0,3,9,9},
      {1,4,2,5,8,7,7,0,0,7,1,2,1,2,7,7,7,4},
      {3,9,7,9,5,8,9,5,6,9,8,8,0,1,4,2,8,2},
      {1,5,2,2,2,5,6,3,9,3,1,7,9,6,8,6,8,3},
      {5,7,8,3,8,8,3,9,9,8,1,9,2,5,4,7,7,7},
      {2,3,2,4,8,5,1,7,2,9,5,2,4,2,9,2,8,7},
      {0,1,6,1,1,0,0,6,5,4,3,4,3,7,9,6,1,9},
	}
	fmt.Println("minPathSum:", minPathSum2(grid))
}
