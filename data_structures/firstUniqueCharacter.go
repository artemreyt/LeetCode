// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
//
// Example 1:
// Input: s = "leetcode"
// Output: 0
//
// Example 2:
// Input: s = "loveleetcode"
// Output: 2
//
// Example 3:
// Input: s = "aabb"
// Output: -1

package datastructures

import "container/list"

type itemFirstUniqChar struct {
	c rune
	i int
}

// queue based on chan
func firstUniqChar(s string) int {
	counterMap := make(map[rune]int)
	queue := make(chan itemFirstUniqChar, len(s))

	for i, c := range s {
		if _, ok := counterMap[c]; !ok {
			counterMap[c] = 0
			queue <- itemFirstUniqChar{c: c, i: i}
		}
		counterMap[c]++
	}

	close(queue)

	for it := range queue {
		if counterMap[it.c] == 1 {
			return it.i
		}
	}
	return -1
}

// queue based on container/list
func firstUniqChar2(s string) int {
	counterMap := make(map[rune]int)
	queue := list.New()

	for i, c := range s {
		if _, ok := counterMap[c]; !ok {
			counterMap[c] = 0
			queue.PushBack(itemFirstUniqChar{c: c, i: i})
		}
		counterMap[c]++
	}

	for queue.Len() > 0 {
		element := queue.Front()
		it := element.Value.(itemFirstUniqChar)
		if counterMap[it.c] == 1 {
			return it.i
		}
		queue.Remove(element)
	}
	return -1
}

// without queue, dumb 2nd string traveral
func firstUniqChar3(s string) int {
	counterMap := make(map[rune]int)

	for _, c := range s {
		if _, ok := counterMap[c]; !ok {
			counterMap[c] = 0
		}
		counterMap[c]++
	}

	for i, c := range s {
		if counterMap[c] == 1 {
			return i
		}
	}
	return -1
}
