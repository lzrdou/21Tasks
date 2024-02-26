package main

import (
	"fmt"
	"io"
)

func main() {
	var nums []int
	var num int
	for {
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}
		nums = append(nums, num)
	}
	fmt.Println(nums)
}
