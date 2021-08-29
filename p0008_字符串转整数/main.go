package main

import (
	"strings"
	"fmt"
	"math"
)

// 解决0000045, - 45 , +1

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	strFix := strings.Split(str, " ")[0]
	strFix = strings.TrimSpace(strFix)
	fmt.Println("strFix1:", strFix, strFix[0])

	if strFix[0] != 45 && strFix[0] != 43 &&  (48 > strFix[0] || strFix[0] > 57) {
		fmt.Println("test", strFix[0])
		return 0
	}
	str = strFix

	for i:=0;i<len(str);i++ {
		if i > 0 {
			if str[i] < 48 || str[i] > 57 {
				str = str[:i]
				fmt.Println(i)
				break
			}
		}
	}

	fmt.Println("str1:", str)
	// 以-开头
	sign := 1
	if str[0] == 45 {
		// 解决-
		if len(str) == 1 {
			return 0
		}
		sign = -1
		str = str[1:]
	}

	if str[0] == 43 {
		// 解决+
		if len(str) == 1 {
			return 0
		}
		str = str[1:]
	}

	var ret int
	for i := 0; i<len(str); i++ {
		
		ret = (ret *10 + int((str[i] - 48))) 

		if ret * sign > math.MaxInt32 {
			return math.MaxInt32
		}

		if  ret * sign< math.MinInt32 {
			return math.MinInt32
		}
	}
	
	return ret*sign
}

func myAtoi2(str string) int {
    res := 0
    i := 0
    
    // get the first non-space character
    for i < len(str) {
        if str[i] == ' ' {
            i++            
        } else {
            break
        }  
    }
    
    sign := 1
    firstNonSpace := true
    for i < len(str) {
        // the  first non-negative value
        if firstNonSpace {
            firstNonSpace = false
            
            if str[i] == '-' || (str[i] >= '0' && str[i] <= '9') || str[i] == '+' {
                if str[i] == '-' {
                    sign = -1
                } else if str[i] >= '0' && str[i] <= '9' {
                    res = res * 10 + int(str[i] - '0')                    
                }
                i++
            } else {
                return res
            }
        } else {
            // if it's not the first non-negative value.
			// we will simple stop when we reach the first non-valid symbol.
            if str[i] >= '0' && str[i] <= '9' {
                res = res * 10 + int(str[i] - '0')
                i++
            } else {
                return res * sign
            }
		}
		
        if res * sign < -2147483648 {
            return -2147483648
        }
        
        if res * sign > 2147483647 {
            return 2147483647
        }
    }

    return res * sign
}

func main() {

	str := "-9223372036854775808aa"
	ret := myAtoi2(str)

	fmt.Println(ret)
}