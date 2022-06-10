package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Dimas",
			request:  "Dimas",
			expected: "hello Dimas",
		},
		{
			name:     "Hartanto",
			request:  "Hartanto",
			expected: "hello Hartanto",
		},
		{
			name:     "Adel",
			request:  "Adel",
			expected: "hello Adel",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWord(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
