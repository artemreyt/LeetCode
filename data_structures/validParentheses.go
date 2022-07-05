// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

package datastructures

var closingToOpeningBracket = map[rune]rune{
	')': '(', ']': '[', '}': '{',
}

type stackRune []rune

func (s *stackRune) Push(val rune) {
	*s = append(*s, val)
}

func (s *stackRune) Pop() {
	if s.Len() > 0 {
		*s = (*s)[:s.Len()-1]
	}
}

func (s *stackRune) Peek() rune {
	if s.Len() == 0 {
		return 0
	}
	return (*s)[s.Len()-1]
}

func (s *stackRune) Len() int {
	return len(*s)
}

func isValid(s string) bool {
	st := make(stackRune, len(s))

	for _, c := range s {
		if opening, ok := closingToOpeningBracket[c]; ok {
			if st.Peek() != opening {
				return false
			}
			st.Pop()
		} else {
			st.Push(c)
		}
	}
	return st.Len() == 0
}
