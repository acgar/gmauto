package rules

import "testing"

func TestCheckTextPattern(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		pattern  string
		expected bool
	}{
		{
			name:     "same text",
			input:    "Hola mundo",
			pattern:  "Hola mundo",
			expected: true,
		},
		{
			name:     "different text",
			input:    "Hola mundo",
			pattern:  "Hola mund",
			expected: false,
		},
		{
			name:     "ending wildcard",
			input:    "hola mundo",
			pattern:  "hola*",
			expected: true,
		},
		{
			name:     "ending wildcard no match",
			input:    "ahola mundo",
			pattern:  "hola*",
			expected: false,
		},
		{
			name:     "starting wildcard",
			input:    "hola mundo",
			pattern:  "*mundo",
			expected: true,
		},
		{
			name:     "starting wildcard no match",
			input:    "hola mundoa",
			pattern:  "*mundo",
			expected: false,
		},
		{
			name:     "middle wildcard",
			input:    "hola mundo",
			pattern:  "hola*mundo",
			expected: true,
		},
		{
			name:     "middle wildcard no match",
			input:    "hola mundoa",
			pattern:  "hola*mundo",
			expected: false,
		},
		{
			name:     "zero chars wildcard",
			input:    "hola mundo",
			pattern:  "*hola mundo",
			expected: true,
		},
		{
			name:     "zero chars wildcard 2",
			input:    "hola mundo",
			pattern:  "hola mundo*",
			expected: true,
		},
		{
			name:     "zero chars wildcard 3",
			input:    "hola mundo",
			pattern:  "hola *mundo",
			expected: true,
		},
		{
			name:     "multiple wildcards",
			input:    "Hola123mundo456nuevo",
			pattern:  "Hola*mundo*nuevo",
			expected: true,
		},
		{
			name:     "multiple wildcards no chars",
			input:    "Holamundonuevo",
			pattern:  "Hola*mundo*nuevo",
			expected: true,
		},
		{
			name:     "last part missing",
			input:    "Hola123mundo",
			pattern:  "Hola*mundo*nuevo",
			expected: false,
		},
		{
			name:     "start part missing",
			input:    "XHola123mundo456nuevo",
			pattern:  "Hola*mundo*nuevo",
			expected: false,
		},
		{
			name:     "wildcard",
			input:    "abc",
			pattern:  "*",
			expected: true,
		},
		{
			name:     "wildcard 2",
			input:    "hola",
			pattern:  "h*a",
			expected: true,
		},
		{
			name:     "double wildcard",
			input:    "Hola mundo",
			pattern:  "Hola**mundo",
			expected: true,
		},
		{
			name:     "triple wildcard",
			input:    "Hola123mundo456nuevo",
			pattern:  "Hola**mundo***nuevo**",
			expected: true,
		},
		{
			name:     "no matching order",
			input:    "Holamundo123nuevo",
			pattern:  "Hola*nuevo*mundo",
			expected: false,
		},
		{
			name:     "no matching",
			input:    "Hola123mundo456nuevo",
			pattern:  "Hola*adios*nuevo",
			expected: false,
		},
		{
			name:     "spaces matters on patterns",
			input:    "Hola que talmundo",
			pattern:  "Hola* mundo",
			expected: false,
		},
		{
			name:     "empty text match wildcard",
			input:    "",
			pattern:  "*",
			expected: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := checkTextPattern(tt.input, tt.pattern)

			if got != tt.expected {
				t.Errorf(
					"checkTextPattern(%q, %q) = %v, want %v",
					tt.input,
					tt.pattern,
					got,
					tt.expected,
				)
			}
		})
	}
}
