package main

import (
	"fmt"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func flatten(list []interface{}) []interface{} {
	res := make([]interface{}, 0)
	stack := []interface{}{list}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curList, ok := cur.([]interface{}); ok {
			for i := len(curList) - 1; i >= 0; i-- {
				stack = append(stack, curList[i])
			}
		} else {
			res = append(res, cur)
		}
	}

	return res
}

func maxSizeSubarraySumEqualsK(array []int, k int) int {
	hashTable := make(map[int]int)
	hashTable[0] = -1
	curSum := 0
	maxSize := 0
	for i := 0; i < len(array); i++ {
		curSum += array[i]
		if index, ok := hashTable[curSum-k]; ok {
			maxSize = max(maxSize, i-index)
		}
		if _, ok := hashTable[curSum]; !ok {
			hashTable[curSum] = i
		}
	}
	return maxSize
}

func allSubsets(arr []interface{}) []interface{} {
	accu := make([]interface{}, 0)
	res := make([]interface{}, 0)
	allSubsetsHelper(arr, 0, accu, &res)
	return res
}

func allSubsetsHelper(arr []interface{}, index int, accu []interface{}, res *[]interface{}) {
	*res = append(*res, accu)
	for i := index; i < len(arr); i++ {
		accu = append(accu, arr[i])
		allSubsetsHelper(arr, i+1, accu, res)
		accu = (accu)[:len(accu)-1]
	}
}

func allSubsetsIterative(arr []int) [][]int {
	res := [][]int{[]int{}}
	for i := 0; i < len(arr); i++ {
		curLength := len(res)
		for j := 0; j < curLength; j++ {
			newSlice := make([]int, len(res[j]))
			copy(newSlice, res[j])
			newSlice = append(newSlice, arr[i])
			res = append(res, newSlice)
		}
	}
	return res
}

func sortTimes(timePoints []string) []string {
	if len(timePoints) <= 1 {
		return timePoints
	}
	mid := len(timePoints) / 2
	return mergeTimes(sortTimes(timePoints[:mid]), sortTimes(timePoints[mid:]))
}

func mergeTimes(timesA, timesB []string) []string {
	i, j := 0, 0
	merged := make([]string, 0)
	for i < len(timesA) && j < len(timesB) {
		timeA := timesA[i]
		timeB := timesB[j]
		if timeA[:2] == timeB[:2] {
			if timeA[3:] <= timeB[3:] {
				merged = append(merged, timeA)
				i++
			} else {
				merged = append(merged, timeB)
				j++
			}
		} else if timeA[:2] < timeB[:2] {
			merged = append(merged, timeA)
			i++
		} else {
			merged = append(merged, timeB)
			j++
		}
	}
	if i < len(timesA) {
		merged = append(merged, timesA[i:]...)
	}
	if j < len(timesB) {
		merged = append(merged, timesB[j:]...)
	}
	return merged
}

func main() {
	list := []interface{}{1, []interface{}{2, []interface{}{3, 4, []interface{}{5, []interface{}{6, 7}, 8}}, 10}}
	res := flatten(list)
	fmt.Printf("flattened list: %v\n", res)

	fmt.Println(maxSizeSubarraySumEqualsK([]int{1, -1, 5, -2, 3}, 3))
	fmt.Println(maxSizeSubarraySumEqualsK([]int{-2, -1, 2, 1}, 1))

	arr := []interface{}{3, 2, 1}
	subsets := allSubsets(arr)
	fmt.Printf("all subsets of %v are: %v\n", arr, subsets)
	arrB := []int{3, 2, 1}
	subsetsB := allSubsetsIterative(arrB)
	fmt.Printf("all subsets of %v are: %v\n", arr, subsetsB)

	timePoints := []string{"23:59", "00:00", "12:00", "04:10"}
	sortedTimes := sortTimes(timePoints)
	fmt.Printf("sorted time of %v is %v\n", timePoints, sortedTimes)

}
