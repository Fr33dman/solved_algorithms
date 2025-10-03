package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/646/

Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.


Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105


Follow up:

Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?
*/

func rotate(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}

	replaces := 0
	spread := 0
	for replaces < len(nums) {
		lastElement := nums[(len(nums)-k%len(nums))+spread]
		for i := 0; i < len(nums) && replaces < len(nums); i++ {
			if (i*k) > 0 && (i*k)%len(nums) == 0 {
				spread++
				break
			}
			nums[(i*k+spread)%len(nums)], lastElement = lastElement, nums[(i*k+spread)%len(nums)]
			replaces++
		}
	}
}

func testRotate() {
	test1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(test1, 3)
	fmt.Println("test1 : ", reflect.DeepEqual(test1, []int{5, 6, 7, 1, 2, 3, 4}))

	test2 := []int{1}
	rotate(test2, 8)
	fmt.Println("test2 : ", reflect.DeepEqual(test2, []int{1}))

	test3 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(test3, 9)
	fmt.Println("test3 : ", reflect.DeepEqual(test3, []int{6, 7, 1, 2, 3, 4, 5}))

	test4 := []int{1, 2, 3, 4, 5, 6}
	rotate(test4, 2)
	fmt.Println("test4 : ", reflect.DeepEqual(test4, []int{5, 6, 1, 2, 3, 4}))

	test5 := []int{1, 2, 3, 4, 5, 6}
	rotate(test5, 9)
	fmt.Println("test5 : ", reflect.DeepEqual(test5, []int{4, 5, 6, 1, 2, 3}))
}
