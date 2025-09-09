package easy

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "III",
			want: 3,
		},
		{
			name: "Example 2",
			s:    "LVIII",
			want: 58,
		},
		{
			name: "Example 3",
			s:    "MCMXC",
			want: 1990,
		},
		{
			name: "Single character",
			s:    "I",
			want: 1,
		},
		{
			name: "Subtraction case IV",
			s:    "IV",
			want: 4,
		},
		{
			name: "Subtraction case IX",
			s:    "IX",
			want: 9,
		},
		{
			name: "Subtraction case XL",
			s:    "XL",
			want: 40,
		},
		{
			name: "Subtraction case XC",
			s:    "XC",
			want: 90,
		},
		{
			name: "Subtraction case CD",
			s:    "CD",
			want: 400,
		},
		{
			name: "Subtraction case CM",
			s:    "CM",
			want: 900,
		},
		{
			name: "Complex number",
			s:    "MCMLIV",
			want: 1954,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInt(tt.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
