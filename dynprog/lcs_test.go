package dyoprog

import (
	"testing"
)

func Test_lengthOfLCS(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{name: "normal1", a: "abcde", b: "aceb", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLCS(tt.a, tt.b); got != tt.want {
				t.Errorf("lengthOfLCS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLCS(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{name: "normal1", a: "abcde", b: "aceb", want: "ace"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLCS(tt.a, tt.b); got != tt.want {
				t.Errorf("findLCS() = %v, want %v", got, tt.want)
			}
		})
	}
}
