package main

import "fmt"

/*
Container With Most Water
https://leetcode.com/problems/container-with-most-water/description/?envType=problem-list-v2&envId=array

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.


Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

type border struct {
	idx    int
	height int
}

func maxArea(height []int) int {
	/*
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	maxVolume := 0
	i, j := 0, len(height)-1
	for i != j {
		volume := min(height[i], height[j]) * (j - i)
		if volume > maxVolume {
			maxVolume = volume
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxVolume
}

func testMaxArea() {
	fmt.Println("test1: ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	fmt.Println("test2: ", maxArea([]int{1, 8, 6, 2, 259, 295, 8, 3, 7}) == 259)
	fmt.Println("test2: ", maxArea([]int{1, 1}) == 1)
}
