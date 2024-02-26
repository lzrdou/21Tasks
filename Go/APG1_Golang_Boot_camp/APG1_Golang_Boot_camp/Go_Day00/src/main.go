package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"sort"
)

type statNums struct {
	length, mode     int
	median, mean, sd float64
	nums             []int
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
	s.mean = float64(sum) / float64(s.length)
}

func (s *statNums) calcMedian() {
	if s.length%2 == 1 {
		s.median = float64(s.nums[s.length/2])
	} else {
		s.median = (float64(s.nums[s.length/2]) + float64(s.nums[s.length/2+1])) / 2
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

func (s *statNums) calcSD() {
	for _, num := range s.nums {
		s.sd += math.Pow(float64(num)-s.mean, 2)
	}
	s.sd = math.Sqrt(s.sd / float64(s.length))
}

func (s *statNums) calculateStats() {
	s.calcMode()
	s.calcMedian()
	s.calcMean()
	s.calcSD()
}

func main() {
	var stat statNums
	args := os.Args

	stat.initStruct()
	stat.calculateStats()

	if slices.Contains(args, "--mean") {
		fmt.Println("Mean:", stat.mean)
	} else if slices.Contains(args, "--median") {
		fmt.Println("Median:", stat.median)
	} else if slices.Contains(args, "--mode") {
		fmt.Println("Mode:", stat.mode)
	} else if slices.Contains(args, "--sd") {
		fmt.Printf("SD: %.2f\n", stat.sd)
	} else if len(args) == 1 {
		fmt.Println("Mean:", stat.mean)
		fmt.Println("Median:", stat.median)
		fmt.Println("Mode:", stat.mode)
		fmt.Printf("SD: %.2f\n", stat.sd)
	}

}
