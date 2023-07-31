package main

import (
	"fmt"
)

/*
	先用最長的長度 -> len(nums) 開始 --
	每次從startIdx + length 開始篩選出nums[startIdx:endIdx] 去sum，
	若有達到result則把length記錄到minLen && shouldBreak = false, 就不用這個length繼續算了，開始下一個length,
	一開始要先把shouldBreak = true, 因為如果沒找到，就代表不用在找了

	reference: https://janac.medium.com/what-is-the-sliding-window-algorithm-f9fcfe92b853
*/

func minSubArrayLen(target int, nums []int) int {
	// 滑動窗口的概念, 選一個窗口(nums[startIdx:endIdx]) 開始滑動比對， 當窗口滑動完了，在加大窗口尺寸，繼續比對
	var (
		minLen      int
		shouldBreak bool
	)
	for length := len(nums); length > 0; length-- {
		shouldBreak = true
		for startIdx := 0; startIdx < len(nums); startIdx++ {
			endIdx := startIdx + length
			if endIdx > len(nums)+1 {
				break
			}
			if result := sum(nums[startIdx:endIdx]...); result >= target {
				minLen = length
				shouldBreak = false
				break
			}
		}
		if minLen != 0 && shouldBreak {
			break
		}
	}

	return minLen
}

func sum(nums ...int) int {
	var output int
	for i := range nums {
		output += nums[i]
	}
	return output
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7

	output := minSubArrayLen(target, nums)
	expected := 2

	if output != expected {
		fmt.Printf("output: %d, expected: %d", output, expected)
	}

	fmt.Println(output)

}
