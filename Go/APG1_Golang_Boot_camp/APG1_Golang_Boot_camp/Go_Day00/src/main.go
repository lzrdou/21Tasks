package main

import (
	"fmt"
	"io"
	"sort"
)

type statNums struct {
	length int
	nums   []int
}

func (s *statNums) InitStruct() {
	num := 0
	for {
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}
		s.nums = append(s.nums, num)
	}
	sort.Slice(s.nums, func(i, j int) bool {
		return s.nums[i] < s.nums[j]
	})
	s.length = len(s.nums)
}

func (s *statNums) getMean() float32 {
	sum := 0
	for _, num := range s.nums {
		sum += num
	}
	return float32(sum) / float32(s.length)
}

func (s *statNums) getMedian() float32 {
	if s.length%2 == 1 {
		return float32(s.nums[s.length/2])
	} else {
		return (float32(s.nums[s.length/2]) + float32(s.nums[s.length/2+1])) / 2
	}
}

func (s *statNums) getMode() int {
	cnt := make(map[int]int)
	for _, num := range s.nums {
		cnt[num] += 1
	}
	mode := 0
	maxFreq := 0
	for _, key := range s.nums {
		freq := cnt[key]
		if freq > maxFreq {
			mode = key
			maxFreq = freq
		}
	}
	return mode
}

func main() {
	var stat statNums

	stat.InitStruct()
	fmt.Println(stat.nums)
	fmt.Println("Mean:", stat.getMean())
	fmt.Println("Median:", stat.getMedian())
	fmt.Println("Mode:", stat.getMode())

}
