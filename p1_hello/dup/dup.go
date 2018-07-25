package main

import(
	"os"
	"fmt"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		countLine(os.Stdin, counts)
	}else{
		for _, farg := range file {
			f, err := os.Open(farg)
			if err != nil {
				fmt.Println("open file error:", err)
				return
			}
			countLine(f, counts)

			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s",n, line)
		}
	}
}

func countLine(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)


	for input.Scan() {
		if input.Text() == "quit" {
			break
		}
		counts[input.Text()]++
	}
	// 忽略input.Err()中错误
}