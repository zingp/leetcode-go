package main

import (
	"fmt"
)


func binarySearch(s []int, val int) bool {
	left := 0
	right := len(s)-1
	for left < right {
		mid := (left+right) / 2
		if s[mid] == val {
			return true
		} else if s[mid] > val {
			right = mid 
		}else {
			left = mid
		}
	}
	return false
}

// 如果是左边小于等于右边下面就要用mid-1或mid+1
func binarySearch2(s[]int, val int) bool {
	left := 0
	right := len(s)-1
	for left <= right {
		mid := (left+right) / 2
		if s[mid] == val {
			return true
		} else if s[mid] > val {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return false
}

func quickSort(array []int, left, right int) {
	if left < right {
		mid := partition(array, left, right)
		quickSort(array, left, mid)
		quickSort(array, mid+1, right)
	}
	
}

func partition(array []int, left, right int) int {
	tmp := array[left]
	for left < right {
		for left<right && array[right] >= tmp {
			right--
		}
		array[left] = array[right]
		for left<right && array[left] <= tmp {
			left++
		}
		array[right] = array[left]
	}
	array[left] = tmp

	return left
}

func main(){
	s := []int {1,2,3,4,5, -1,-5,-10}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}