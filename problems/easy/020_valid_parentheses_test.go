package easy

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "()",
			want: true,
		},
		{
			name: "Example 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Example 3",
			s:    "(]",
			want: false,
		},
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
		{
			name: "Nested parentheses",
			s:    "((()))",
			want: true,
		},
		{
			name: "Mixed nested",
			s:    "([{}])",
			want: true,
		},
		{
			name: "Wrong order",
			s:    "([)]",
			want: false,
		},
		{
			name: "Only opening",
			s:    "(((",
			want: false,
		},
		{
			name: "Only closing",
			s:    ")))",
			want: false,
		},
		{
			name: "Mismatched",
			s:    "(}",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
