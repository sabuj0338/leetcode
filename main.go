package main

import (
	"fmt"
	p1 "leetcode/p1"
)

func main() {
	var nums = []int{3, 2, 3}
	var target int = 6

	var result = p1.TwoSum(nums, target)
	fmt.Println(result)
}
