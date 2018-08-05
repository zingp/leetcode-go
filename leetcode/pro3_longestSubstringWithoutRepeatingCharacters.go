package main

import (
	"fmt"
)
// 最长没有重复字符的子串，返回长度

func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
    var bsMaxlen, start int
    bucket := map[byte]int{}
    for i := 0; i < len(bs); i++ {
        if _, ok := bucket[bs[i]]; ok && start <= bucket[bs[i]]{
            start = bucket[bs[i]] + 1
        }else{
            if bsMaxlen < i - start + 1 {
                bsMaxlen = i - start + 1
            }
        }
        bucket[bs[i]] = i
    }
    return bsMaxlen
}

func main() {
	s1 := "aaaaa"
	s2 := "abcab"
	s3 := "pwwkew"
	s4 := "dvdf"

	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println(lengthOfLongestSubstring(s4))
}

//2
func lengthOfLongestSubstring2(s string) int {
    if len(s) < 2 {
        return len(s)
    }
    length := len(s)
    i, j := 0, 1
    max := 0
    for j < length {
        index, ok := contain(s, i, j)
        if ok {
            if newLen := j - i; newLen > max {
                max = newLen
            }
            i = index + 1
        } 
        j++
    }
    if newLen := j - i; newLen > max {
        max = newLen
    }
    return max
}

func contain(s string, i, j int)(int, bool) {
    str := s[i:j]
    for index, v := range str {
        if byte(v) == s[j] {
            return index + i, true
        }
    }
    return 0, false
}