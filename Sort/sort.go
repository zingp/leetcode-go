package main

import(
	"fmt"
)

func bubbleSort(array []int) []int{

	for i:=0;i<len(array);i++ {
		for j:=0;j<(len(array)-i-1);j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] =  array[j+1], array[j]
			}
		}
	}
	return array
}

func bubbleSort2(array []int) []int {
	for i:=0;i<len(array);i++ {
		flag := true
		for j:=0;j<(len(array)-i-1);j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] =  array[j+1], array[j]
				flag = false
			}
		}
		if flag {
			return array
		}
	}
	return array
}

func insertSort(array []int) []int {
	for i:=1;i<len(array);i++ {
		tmp := array[i]
		j := i-1
		for j>=0 && array[j] > tmp{
			array[j+1] = array[j]
			j--
		}
		array[j+1] = tmp
	}
	return array
}

func selectSort(array []int) []int{
	for i:=0;i<len(array);i++ {
		minIndex := i
		for j:=i+1;j<len(array);j++{
			if array[i] > array[j]{
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex],array[i]
	}
	return array
}

func quickSort(array []int, left int, right int) {
	if left < right {
		mid := partition(array, left, right)
		quickSort(array, left, mid)
		quickSort(array, mid+1, right)
	}
}

func partition(array []int, left int, right int) int{
	tmp := array[left]
	for left < right {
		for left < right && array[right] >= tmp {
            right--
		}
		array[left] = array[right]
		for left < right && array[left] <= tmp{
			left++
		}		
		array[right] = array[left]	
	}
	array[left] = tmp
	return left
}

func heapSort(array []int) {
	n := len(array)
	for i:=(n/2-1);i>=0;i--{
		sift(array, i, n-1)
	}
	for i:=n-1;i>=0;i--{
		array[0],array[i] = array[i], array[0]
		sift(array,0,i-1)
	}
}


func sift(array []int, left, right int) {
	i := left
	j := 2*i+1
	tmp := array[i]
	for j <= right{
		if j < right && array[j]< array[j+1] {
			j++
		}
		if array[j] > tmp {
			array[i] = array[j]
			i = j
			j = 2*i+1
		}else {
			break
		}
	}
	array[i] = tmp
}

func main() {
	arr := []int {9, 8, 7,3,6,1,-1}
	heapSort(arr)
	fmt.Println(arr)
}