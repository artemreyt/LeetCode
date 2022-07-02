// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true
//
// Example 2:
// Input: s = "rat", t = "car"
// Output: false
//
// Constraints: s and t consist of lowercase English letters.

package datastructures

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var stock [26]int
	for _, c := range s {
		stock[c-'a']++
	}

	for _, c := range t {
		if stock[c-'a'] <= 0 {
			return false
		}
		stock[c-'a']--
	}
	return true
}
