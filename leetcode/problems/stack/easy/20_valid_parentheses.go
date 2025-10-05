package main

import "fmt"

/*
Valid Parentheses
https://leetcode.com/problems/valid-parentheses/description/?envType=problem-list-v2&envId=stack

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true

Example 5:

Input: s = "([)]"

Output: false



Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
	/*
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	if len(s) == 0 {
		return true
	} else if len(s)%2 != 0 {
		return false
	}

	stack := make([]int32, 0, len(s)/2)
	for _, r := range s {
		switch r {
		case 40, 91, 123:
			stack = append(stack, r)
		case 93, 125:
			if len(stack) == 0 || stack[len(stack)-1] != r-2 {
				return false
			}
			stack = stack[:len(stack)-1]
		case 41:
			if len(stack) == 0 || stack[len(stack)-1] != 40 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func testIsValid() {
	fmt.Println("test1: ", isValid("[{()}]") == true)
	fmt.Println("test2: ", isValid("][") == false)
	fmt.Println("test2: ", isValid("()") == true)
	fmt.Println("test2: ", isValid("([]{[]}())") == true)
	fmt.Println("test2: ", isValid("([]{[)}())") == false)
}
