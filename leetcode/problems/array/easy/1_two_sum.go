package main

import (
	"fmt"
	"reflect"
)

/*
Two Sum
https://leetcode.com/problems/two-sum/?envType=problem-list-v2&envId=array

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

type testFixture struct {
	data   []int
	target int
	answer []int
}

// First solution (oh I'm so dumb...)
type item struct {
	idx int
	val int
}

func mySearch(nums []item, min, target int) int {
	lft, rgt := 0, len(nums)-1
	idx := len(nums) / 2
	for {
		if rgt-lft <= 1 {
			if nums[rgt].val+min <= target {
				return rgt
			}
			return lft
		}
		if nums[idx].val+min == target {
			return idx
		} else if nums[idx].val+min < target {
			lft, idx = idx, (rgt+idx)/2
		} else {
			rgt, idx = idx, (idx+lft)/2
		}
	}
}

func mySort(idx int, nums []int) []item {
	sortArray := func(arr1, arr2 []item) []item {
		i, j := 0, 0
		res := make([]item, 0)
		for i < len(arr1) && j < len(arr2) {
			if arr1[i].val < arr2[j].val {
				res = append(res, arr1[i])
				i++
			} else {
				res = append(res, arr2[j])
				j++
			}
		}
		res = append(res, arr1[i:]...)
		res = append(res, arr2[j:]...)
		return res
	}

	if len(nums) == 0 {
		return []item{}
	} else if len(nums) == 1 {
		return []item{{idx: idx, val: nums[0]}}
	} else if len(nums) == 2 {
		if nums[0] < nums[1] {
			return []item{{idx: idx, val: nums[0]}, {idx: idx + 1, val: nums[1]}}
		} else {
			return []item{{idx: idx + 1, val: nums[1]}, {idx: idx, val: nums[0]}}
		}
	}

	mid := len(nums) / 2
	arr1 := mySort(idx, nums[:mid])
	arr2 := mySort(idx+mid, nums[mid:])
	return sortArray(arr1, arr2)
}

func twoSum(nums []int, target int) []int {
	/*
		Time complexity: O(n log(n))
		Space complexity: O(n)
	*/
	itemNums := mySort(0, nums)
	itemNums = itemNums[:mySearch(itemNums, itemNums[0].val, target)+1]
	i, j := 0, len(itemNums)-1
	for {
		if i == j {
			return []int{}
		}
		if itemNums[i].val+itemNums[j].val == target {
			if itemNums[i].idx < itemNums[j].idx {
				return []int{itemNums[i].idx, itemNums[j].idx}
			} else {
				return []int{itemNums[j].idx, itemNums[i].idx}
			}
		} else if itemNums[i].val+itemNums[j].val < target {
			i++
		} else {
			j--
		}
	}
}

// Second solution (final)
func twoSumV2(nums []int, target int) []int {
	/*
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	hashMap := make(map[int]int)

	for i, val := range nums {
		if idx, isFound := hashMap[target-val]; isFound {
			if i < idx {
				return []int{i, idx}
			}
			return []int{idx, i}
		}
		hashMap[val] = i
	}

	return []int{}
}

func testTwoSum() {
	testData := []testFixture{
		{[]int{2, 7, 9, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 5, 3, 1, 7, 6, 11, 15, 4}, 9, []int{0, 4}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
	}
	for i, fixture := range testData {
		fmt.Printf("test%d : %v\n", i, reflect.DeepEqual(twoSum(fixture.data, fixture.target), fixture.answer))
	}
	for i, fixture := range testData {
		fmt.Printf("V2 test%d : %v\n", i, reflect.DeepEqual(twoSumV2(fixture.data, fixture.target), fixture.answer))
	}
}
