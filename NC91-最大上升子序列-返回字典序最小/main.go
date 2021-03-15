package main

import (
	"fmt"
)

func LIS(arr []int) []int {

	dp := make([]int, len(arr))
	dp[0] = 1
	ans := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i] > ans[len(ans)-1] {
			ans = append(ans, arr[i])
			dp[i] = len(ans)
		} else {
			left, right := 0, len(ans)-1
			for left < right {
				mid := left + (right-left)>>1
				if ans[mid] >= arr[i] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			ans[left] = arr[i]
			dp[i] = left + 1
		}
	}
	for i, j := len(dp)-1, len(ans); j > 0; i-- {
		if dp[i] == j {
			j--
			ans[j] = arr[i]
		}
	}

	return ans
}

func main() {
	arr := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	ans := LIS(arr)
	fmt.Println(ans)
}
