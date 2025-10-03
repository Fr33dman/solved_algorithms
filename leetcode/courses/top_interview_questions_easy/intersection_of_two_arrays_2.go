package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

/*
Intersection of Two Arrays II
https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/674/

Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.



Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.


Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000


Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

func intersect(nums1 []int, nums2 []int) []int {

	sliceWithDefaults := func(num int, length int) []int {
		res := make([]int, length)
		for i := 0; i < length; i++ {
			res[i] = num
		}
		return res
	}

	count := func(nums []int) [1001]int {
		numsCounter := [1001]int{}
		for _, num := range nums {
			numsCounter[num]++
		}
		return numsCounter
	}

	count1, count2 := count(nums1), count(nums2)

	resSlice := make([]int, 0)

	for i := 0; i < 1001; i++ {
		resSlice = append(resSlice, sliceWithDefaults(i, min(count1[i], count2[i]))...)
	}

	return resSlice
}

func intersectAsync(nums1 []int, nums2 []int) []int {

	sliceWithDefaults := func(num int, length int) []int {
		res := make([]int, length)
		for i := 0; i < length; i++ {
			res[i] = num
		}
		return res
	}

	count := func(nums []int, wg *sync.WaitGroup, resCh chan [1001]int) {
		numsCounter := [1001]int{}
		for _, num := range nums {
			numsCounter[num]++
		}
		wg.Done()
		resCh <- numsCounter
	}

	resCh := make(chan [1001]int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go count(nums1, wg, resCh)
	go count(nums2, wg, resCh)
	wg.Wait()

	count1, count2 := <-resCh, <-resCh
	close(resCh)

	resSlice := make([]int, 0)

	for i := 0; i < 1001; i++ {
		resSlice = append(resSlice, sliceWithDefaults(i, min(count1[i], count2[i]))...)
	}

	return resSlice
}

func generateSlice() []int {
	res := make([]int, rand.Int()%int(math.Pow(10, 7)))
	for i := 0; i < len(res); i++ {
		res[i] = rand.Int() % 1000
	}
	return res
}

func testIntersect() {
	fmt.Println("test1: ", reflect.DeepEqual(intersect([]int{1, 2, 2, 1}, []int{2, 2}), []int{2, 2}))
	fmt.Println("test2: ", reflect.DeepEqual(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{4, 9}))

	// Benchmark
	benchArrays := make(map[int][][]int)
	for i := range 10 {
		benchArrays[i] = [][]int{generateSlice(), generateSlice()}
	}

	start := time.Now()
	for i := range 10 {
		intersect(benchArrays[i][0], benchArrays[i][1])
		fmt.Printf("bench%d: done\n", i)
	}
	since := time.Since(start)
	fmt.Println("elapsed time async: ", since)

	start = time.Now()
	for i := range 10 {
		intersectAsync(benchArrays[i][0], benchArrays[i][1])
		fmt.Printf("bench%d async: done\n", i)
	}
	since = time.Since(start)

	fmt.Println("elapsed time async: ", since)
}
