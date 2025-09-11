package dyoprog

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected int
	}{
		{
			name:     "both empty strings",
			word1:    "",
			word2:    "",
			expected: 0,
		},
		{
			name:     "first string empty",
			word1:    "",
			word2:    "abc",
			expected: 3,
		},
		{
			name:     "second string empty",
			word1:    "abc",
			word2:    "",
			expected: 3,
		},
		{
			name:     "identical strings",
			word1:    "kitten",
			word2:    "kitten",
			expected: 0,
		},
		{
			name:     "classic example: kitten -> sitting",
			word1:    "kitten",
			word2:    "sitting",
			expected: 3,
		},
		{
			name:     "single character difference",
			word1:    "abc",
			word2:    "abd",
			expected: 1,
		},
		{
			name:     "insertion only",
			word1:    "abc",
			word2:    "abcd",
			expected: 1,
		},
		{
			name:     "deletion only",
			word1:    "abcd",
			word2:    "abc",
			expected: 1,
		},
		{
			name:     "substitution only",
			word1:    "abc",
			word2:    "axc",
			expected: 1,
		},
		{
			name:     "multiple operations",
			word1:    "horse",
			word2:    "ros",
			expected: 3,
		},
		{
			name:     "intention -> execution",
			word1:    "intention",
			word2:    "execution",
			expected: 5,
		},
		{
			name:     "completely different strings",
			word1:    "abc",
			word2:    "xyz",
			expected: 3,
		},
		{
			name:     "with special characters",
			word1:    "hello world",
			word2:    "hello world!",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDistance2(tt.word1, tt.word2)
			if result != tt.expected {
				t.Errorf("minDistance(%q, %q) = %d, expected %d",
					tt.word1, tt.word2, result, tt.expected)
			}
		})
	}
}

func TestMinDistanceSymmetry(t *testing.T) {
	// 测试编辑距离的对称性：dist(a,b) == dist(b,a)
	testCases := []struct {
		a string
		b string
	}{
		{"kitten", "sitting"},
		{"horse", "ros"},
		{"abc", "xyz"},
		{"", "hello"},
		{"intention", "execution"},
	}

	for _, tc := range testCases {
		dist1 := minDistance(tc.a, tc.b)
		dist2 := minDistance(tc.b, tc.a)
		if dist1 != dist2 {
			t.Errorf("Asymmetry detected: minDistance(%q, %q) = %d, minDistance(%q, %q) = %d",
				tc.a, tc.b, dist1, tc.b, tc.a, dist2)
		}
	}
}
