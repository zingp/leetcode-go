package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverseNum := 0 
	n := x
	for n!= 0 {
		reverseNum = reverseNum*10 + n % 10
		n /= 10
	}

	if reverseNum == x {
		return true
	}
	return false
}

func main() {
	v :=  0
	fmt.Println(isPalindrome(v))
}