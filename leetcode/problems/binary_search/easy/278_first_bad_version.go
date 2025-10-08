package main

import "fmt"

/*
First Bad Version
https://leetcode.com/problems/first-bad-version/description/?envType=problem-list-v2&envId=binary-search

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.



Example 1:

Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.
Example 2:

Input: n = 1, bad = 1
Output: 1


Constraints:

1 <= bad <= n <= 231 - 1
*/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	if version >= 1 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	/*
		Time complexity: O(log(n))
		Space complexity: O(1)
	*/
	lft, i, rgt := 1, n/2, n
	for {
		if rgt-lft <= 1 {
			return rgt
		}
		if isBadVersion(i) {
			rgt, i = i, (rgt+lft)/2
		} else {
			lft, i = i, (i+rgt)/2
		}
	}
}

func testFirstBadVersion() {
	fmt.Println("test1: ", firstBadVersion(1))
}
