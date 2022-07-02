package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstUniqChar(t *testing.T) {
	require.Equal(t, 0, firstUniqChar("leetcode"))
	require.Equal(t, 4, firstUniqChar("leetcodeleet"))
	require.Equal(t, -1, firstUniqChar("leetcodeleetcode"))
}
