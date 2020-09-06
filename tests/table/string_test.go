package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - indices: expected (%d) <> encounter (%d)."

func TestIndex(t *testing.T) {
	//Creating dataset to test some function in string package from Go.
	tests := []struct{
		text string
		part string
		expected int
	}{
		{"Alice is cool", "Alice", 0},
		{"", "", 0},
		{"Ops", "o", -1},
		{"leonard", "o", 2},
	}
	for _, test := range tests{
		t.Logf("Mass: %v", test)
		actual := strings.Index(test.text, test.part)
		if actual != test.expected{
			t.Errorf(msgIndex, test.text, test.part, test.expected, actual)
		}
	}
}

