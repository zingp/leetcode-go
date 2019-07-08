package main
/*
最大子序列和问题
https://leetcode.com/problems/maximum-subarray/

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/
import(
	"fmt"
)

// 分治

// 动态规划

// 在线算法
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	for i:= 0; i<len(nums);i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum = curSum + nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func main() {
	// array := []int{-2,1,-3,4,-1,2,1,-5,4}
	array := []int{-2,-3}
	fmt.Println(maxSubArray(array))
}
