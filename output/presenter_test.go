package output

import (
	"testing"
)

func TestPrepare(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "Test with two simple strings",
			input: []string{"abc", "def"},
			want:  "abc def",
		},
		{
			name:  "Test with two empty strings",
			input: []string{"", ""},
			want:  " ",
		},
		{
			name:  "Test with many short strings",
			input: []string{"a", "b", "c", "d", "e", "f"},
			want:  "a b c d e f",
		},
		{
			name:  "Test with many simple strings",
			input: []string{"abc", "bcd", "cde", "def", "efg", "fgh"},
			want:  "abc bcd cde def efg fgh",
		},
		{
			name:  "Test with two address strings",
			input: []string{"http://www.company.com", "http://google.com"},
			want:  "http://www.company.com http://google.com",
		},
		{
			name: "Test with many address strings",
			input: []string{
				"http://google.com",
				"https://github.com",
				"https://stackoverflow.com",
				"https://www.linkedin.com",
				"https://www.facebook.com",
				"https://www.the-scorpions.com",
				"https://www.nightwish.com",
				"https://twitter.com",
				"http://www.astronio.gr",
				"https://felifromgermany.com",
				"https://www.linkedin.com",
			},
			want: "http://google.com https://github.com https://stackoverflow.com https://www.linkedin.com https://www.facebook.com https://www.the-scorpions.com https://www.nightwish.com https://twitter.com http://www.astronio.gr https://felifromgermany.com https://www.linkedin.com",
		},
	}

	presenter := NewPresenter()

	for _, tt := range tests {
		got := presenter.Prepare(tt.input)
		if tt.want != got {
			t.Errorf("Test \"%s\" failed for Input: %+v, Expected Output: %s, Actual Output: %s", tt.name, tt.input, tt.want, got)
		}
	}
}
