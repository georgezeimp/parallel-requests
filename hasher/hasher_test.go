package hasher

import (
	"testing"
)

func TestToMD5(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "test",
			want:  "098f6bcd4621d373cade4e832627b4f6",
		},
		{
			input: "http://google.com",
			want:  "c7b920f57e553df2bb68272f61570210",
		},
		{
			input: "\\m/",
			want:  "25d2222512f85662dd7828fd0be3c86f",
		},
		{
			input: "",
			want:  "d41d8cd98f00b204e9800998ecf8427e",
		},
	}

	hasher := NewHasher()

	for _, tt := range tests {
		got := hasher.ToMD5([]byte(tt.input))
		if tt.want != got {
			t.Errorf("Test failed for Input: %+v, Expected Output: %s, Actual Output: %s", tt.input, tt.want, got)
		}
	}
}
