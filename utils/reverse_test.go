package utils

import "testing"

func TestBytesReverse(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{
			input: []byte("Hello world"),
			want:  "dlrow olleH",
		},
		{
			input: []byte("World hello"),
			want:  "olleh dlroW",
		},
	}

	for _, test := range tests {
		if got := BytesReverse(test.input); string(got) != test.want {
			t.Errorf("BytesReverse(%q) got: %s, want: %s", test.input, got, test.want)
		}
	}
}
