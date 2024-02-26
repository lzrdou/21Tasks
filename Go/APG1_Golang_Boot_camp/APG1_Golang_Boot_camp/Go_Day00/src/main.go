package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	var nums []int
	var num, mode int
	var median float32

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

	l := len(nums)

	cnt := make(map[int]int)
	sum := 0
	for _, num := range nums {
		cnt[num] += 1
		sum += num
	}
	fmt.Printf("Mean: %f\n", float32(sum)/float32(l))

	if l%2 == 1 {
		median = float32(nums[l/2])
	} else {
		median = (float32(nums[l/2]) + float32(nums[l/2+1])) / 2
	}

	maxFreq := 0
	for _, key := range nums {
		freq := cnt[key]
		if freq > maxFreq {
			mode = key
			maxFreq = freq
		}
	}
	fmt.Printf("Median: %f\n", median)
	fmt.Printf("Mode %d\n", mode)

}
