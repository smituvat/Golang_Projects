package main

/*
Given an array of integers, find the length of the longest sub-sequence such that elements in the subsequence are consecutive integers,
 the consecutive numbers can be in any order.

Examples:

Input: arr[] = {1, 9, 3, 10, 4, 20, 2}
Output: 4
Explanation: The subsequence 1, 3, 4, 2 is the longest subsequence of consecutive elements

Input: arr[] = {36, 41, 56, 35, 44, 33, 34, 92, 43, 32, 42}
Output: 5
Explanation: The subsequence 36, 35, 33, 34, 32 is the longest subsequence of consecutive elements.
*/

// func subSeq(arr []int) int {
// 	ans := sort(arr)
// 	newAns := []int{}
// 	for i := 1; i < len(ans)-1; i++ {
// 		if ans[i] == ans[i-1] + 1 {
// 			newAns = appned(newAns, ans[i])
// 		} else {
// 			if max > len(newAns) {
// 				max := len(newAns)
// 			}
// 			newAns = newAns(0:)
// 		}
// 	}
// 	return max
// }
