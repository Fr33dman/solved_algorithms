package main

import "fmt"

/*
Median of Two Sorted Arrays
https://leetcode.com/problems/median-of-two-sorted-arrays/description/?envType=problem-list-v2&envId=array

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).


Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.


Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10^6 <= nums1[i], nums2[i] <= 10^6
*/

func splitArray(nums []int, n int) int {
	lft, idx, rgt := 0, len(nums)/2, len(nums)-1
	for {
		if rgt-lft == 1 {
			if nums[lft] <= n && n <= nums[rgt] {
				return lft
			}
			return rgt
		}

		if nums[idx] < n {
			lft = idx
			idx = (rgt + lft) / 2
		} else if nums[idx] > n {
			rgt = idx
			idx = (rgt + lft) / 2
		} else {
			return idx
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/*
		I've tried to solve it with complexity O((m+n) * log(n))
		which is a bit faster than regular median(sorted(arr1 + arr2))
		but not satisfies problem condition of complexity O(log(m+n))

		I get median of bigger array, then I split second array for 2 parts, to the left and right sides of
		median of first array. Then I find the shift of the median (to left or right side) and
		iterate over the required parts of the arrays and find the median by enumeration, knowing its offset
	*/

	if len(nums1) < len(nums2) {
		nums2, nums1 = nums1, nums2
	}

	mid := nums1[len(nums1)/2]
	border := splitArray(nums2, mid)
	lft := len(nums1)/2 + border + 1
	rgt := (len(nums1) - len(nums1)/2) + (len(nums2) - border - 1)

	fmt.Println(border)
	fmt.Println(lft, rgt)

	i, j := 0, 0
	var lst int
	if lft == rgt {
		return float64(mid)
	} else if lft < rgt {
		end := false
		for {
			if i+j >= (rgt-lft)/2 {
				if (len(nums1)+len(nums2))%2 == 0 {
					if !end {
						end = true
						goto lop
					}
					return float64(min(nums1[len(nums1)/2+i], nums2[border+j])+lst) / 2
				}
				return float64(min(nums1[len(nums1)/2+i], nums2[border+j]))
			}
		lop:
			lst = min(nums1[len(nums1)/2+i], nums2[border+j])
			if len(nums1)/2+i > len(nums1) {
				j++
				continue
			} else if border+j > len(nums2) {
				i++
				continue
			}
			if nums1[len(nums1)/2+i] < nums2[border+j] {
				i++
			} else {
				j++
			}
		}
	} else {
		for {
			if i+j >= (rgt-lft)/2 {
				if (len(nums1)+len(nums2))%2 == 0 {
					return float64(max(nums1[len(nums1)/2-i], nums2[border-j])+lst) / 2
				}
				return float64(max(nums1[len(nums1)/2-i], nums2[border-j]))
			}
			lst = max(nums1[len(nums1)/2-i], nums2[border-j])
			if len(nums1)/2-i < 0 {
				j++
				continue
			} else if border-j < 0 {
				i++
				continue
			}
			if nums1[len(nums1)/2-i] > nums2[border-j] {
				i++
			} else {
				j++
			}
		}
	}
}

func testFindMedianSortedArrays() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4, 6, 8, 10, 11}, []int{5, 7, 9, 12}))
	fmt.Println(findMedianSortedArrays([]int{5, 7, 9, 12}, []int{1, 2, 3, 4, 6, 8, 10, 11}))
}
