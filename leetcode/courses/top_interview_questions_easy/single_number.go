package main

import "fmt"

/*
Single Number
https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/549/

Solution
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.


Example 1:

Input: nums = [2,2,1]

Output: 1

Example 2:

Input: nums = [4,1,2,1,2]

Output: 4

Example 3:

Input: nums = [1]

Output: 1



Constraints:

1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/

func singleNumber(nums []int) int {
	appears := make(map[int]interface{})
	for _, num := range nums {
		if _, ok := appears[num]; ok {
			delete(appears, num)
		} else {
			appears[num] = nil
		}
	}
	for k, _ := range appears {
		return k
	}
	return 0
}

func singleNumberV2(nums []int) int {
	// I admit honestly, this solution is not mine, I didn’t think of it, but I left it here so I can come back to it in the future
	res := 0
	for _, n := range nums {
		//fmt.Printf("%05b %05b\n", res, n)
		res ^= n
	}
	return res
}

func testSingleNumber() {
	fmt.Println("test1: ", singleNumber([]int{2, 2, 1}) == 1)
	fmt.Println("test2: ", singleNumber([]int{4, 1, 2, 1, 2}) == 4)
	fmt.Println("test3: ", singleNumber([]int{1}) == 1)

	fmt.Println("test4: ", singleNumberV2([]int{2, 2, 1}) == 1)
	fmt.Println("test5: ", singleNumberV2([]int{4, 1, 5, 2, 5, 1, 2}) == 4)
	fmt.Println("test6: ", singleNumberV2([]int{1}) == 1)
}
