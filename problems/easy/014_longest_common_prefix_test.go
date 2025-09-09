package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "Example 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Example 2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "Empty array",
			strs: []string{},
			want: "",
		},
		{
			name: "Single string",
			strs: []string{"single"},
			want: "single",
		},
		{
			name: "All same strings",
			strs: []string{"test", "test", "test"},
			want: "test",
		},
		{
			name: "One empty string",
			strs: []string{"abc", "", "ab"},
			want: "",
		},
		{
			name: "No common prefix",
			strs: []string{"abc", "def", "ghi"},
			want: "",
		},
		{
			name: "Partial match",
			strs: []string{"interspecies", "interstellar", "interstate"},
			want: "inters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
