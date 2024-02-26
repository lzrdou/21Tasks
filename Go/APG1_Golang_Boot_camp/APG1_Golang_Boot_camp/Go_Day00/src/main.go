package main

import (
	"fmt"
	"io"
	"sort"
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
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)

}
