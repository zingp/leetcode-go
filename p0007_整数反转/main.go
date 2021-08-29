package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
	"math/rand"
)

// 方法一：通过转换成字符串，反转，再转换成int
func reverse(x int) int {
	intStr := strconv.Itoa(x)

	if len(intStr) == 1 {
		return x
	}

	var strSlice []string
	if intStr[0] == 45 {
		strSlice = append(strSlice, string(intStr[0]))
		intStr = intStr[1:]
	}
	if intStr[len(intStr)-1] == 48 {
		intStr = intStr[:len(intStr)-1]
	}

	for k:=0; k<len(intStr);k++ {
		strSlice = append(strSlice, string(intStr[len(intStr)-k-1]))
	}

	var str string
	str = strings.Join(strSlice, "")
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("strconv string to int error:%v", err)
	}

	// 溢出时返回0
	if value > math.MaxInt32 || value < math.MinInt32 {
        return 0
    }
	
	return value
}

// 方法二：对x取余，获得最后一位数，x取商。反转后的数字等于前一次余数乘以10加当前余数
func reverse2(x int) int {
	var ret int
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}

	if ret > math.MaxInt32 || ret < math.MinInt32 {
        return 0
	}
	
	return ret
}

func main()  {
	for i := 0; i <10000;i++ {
		x := rand.Intn(2 * math.MaxInt32) - math.MaxInt32
		r1 := reverse(x)
		r2 := reverse2(x)
		if r1 == r2 {
			fmt.Printf("%d -- %d\n", r1, r2)
		}
	}
}