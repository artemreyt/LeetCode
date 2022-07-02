// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.
//
// Example 1:
// Input: ransomNote = "a", magazine = "b"
// Output: false
//
// Example 2:
// Input: ransomNote = "aa", magazine = "ab"
// Output: false
//
// Example 3:
// Input: ransomNote = "aa", magazine = "aab"
// Output: true

package datastructures

func canConstruct(ransomNote string, magazine string) bool {
	// Note: we could use array or even a bit sequence instead of map
	// but a map seems to be a more general solution
	availableChars := make(map[rune]int)
	for _, c := range magazine {
		if _, ok := availableChars[c]; !ok {
			availableChars[c] = 0
		}
		availableChars[c]++
	}

	for _, c := range ransomNote {
		if stock, ok := availableChars[c]; ok && stock > 0 {
			availableChars[c]--
		} else {
			return false
		}
	}
	return true
}
