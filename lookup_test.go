package caa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextInHierarchy(t *testing.T) {
	tests := []struct {
		TestName string
		Addr     string
		Name     string
		Rest     string
	}{
		{"Empty", "", "", ""},
		{"Root", ".", ".", ""},
		{"Single", "org.", "org", "."},
		{"Multuple", "example.org.", "example", "org."},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			name, rest := nextInHierarchy(test.Addr)
			assert.Equal(t, test.Name, name)
			assert.Equal(t, test.Rest, rest)
		})
	}
}
