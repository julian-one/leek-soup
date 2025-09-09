package easy

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "Example 1",
			x:    121,
			want: true,
		},
		{
			name: "Example 2",
			x:    -121,
			want: false,
		},
		{
			name: "Example 3",
			x:    10,
			want: false,
		},
		{
			name: "Single digit",
			x:    0,
			want: true,
		},
		{
			name: "Single digit positive",
			x:    7,
			want: true,
		},
		{
			name: "Even length palindrome",
			x:    1221,
			want: true,
		},
		{
			name: "Large palindrome",
			x:    12321,
			want: true,
		},
		{
			name: "Not palindrome",
			x:    123,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.x); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
