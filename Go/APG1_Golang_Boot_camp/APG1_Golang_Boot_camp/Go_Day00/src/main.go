package main

import (
	"fmt"
	"io"
	"sort"
)

type statNums struct {
	length, mode int
	median, mean float32
	nums         []int
}

func (s *statNums) initStruct() {
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

func (s *statNums) calcMean() {
	sum := 0
	for _, num := range s.nums {
		sum += num
	}
	s.mean = float32(sum) / float32(s.length)
}

func (s *statNums) calcMedian() {
	if s.length%2 == 1 {
		s.median = float32(s.nums[s.length/2])
	} else {
		s.median = (float32(s.nums[s.length/2]) + float32(s.nums[s.length/2+1])) / 2
	}
}

func (s *statNums) calcMode() {
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
	s.mode = mode
}

func (s *statNums) calculateStats() {
	s.calcMode()
	s.calcMedian()
	s.calcMean()
}

func main() {
	var stat statNums

	stat.initStruct()
	stat.calculateStats()

	fmt.Println("Mean:", stat.mean)
	fmt.Println("Median:", stat.median)
	fmt.Println("Mode:", stat.mode)

}
