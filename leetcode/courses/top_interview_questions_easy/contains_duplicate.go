package main

import "fmt"

/*
https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/578/

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.


Example 1:

Input: nums = [1,2,3,1]

Output: true

Explanation:

The element 1 occurs at the indices 0 and 3.

Example 2:

Input: nums = [1,2,3,4]

Output: false

Explanation:

All elements are distinct.

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]

Output: true


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

func containsDuplicate(nums []int) bool {
	appeared := make(map[int]interface{})
	for _, n := range nums {
		if _, ok := appeared[n]; ok {
			return true
		}
		appeared[n] = nil
	}
	return false
}

func testContainsDuplicate() {
	fmt.Println("test1: ", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) == true)
	fmt.Println("test1: ", containsDuplicate([]int{1, 2, 3, 4}) == false)
	fmt.Println("test1: ", containsDuplicate([]int{1, 3, 2, 5, 4, 8, 6, 9, 7, 10}) == false)
	fmt.Println("test1: ", containsDuplicate([]int{1, 3, 2, 5, 4, 8, 6, 9, 7, 10, 5}) == true)
}
