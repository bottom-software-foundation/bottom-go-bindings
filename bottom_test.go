package bottom

import (
	"testing"
)

type testCase struct {
	value    string
	expected string
}

var testCases [4]testCase = [4]testCase{
	{"Test", "💖✨✨✨,,,,👉👈💖💖,👉👈💖💖✨🥺👉👈💖💖✨🥺,👉👈"},
	{"h", "💖💖,,,,👉👈"},
	{"🥺", "🫂✨✨✨✨👉👈💖💖💖🥺,,,,👉👈💖💖💖✨🥺👉👈💖💖💖✨✨✨🥺,👉👈"},
	{"がんばれ", "🫂✨✨🥺,,👉👈💖💖✨✨🥺,,,,👉👈💖💖✨✨✨✨👉👈🫂✨✨🥺,,👉👈💖💖✨✨✨👉👈💖💖✨✨✨✨🥺,,👉👈🫂✨✨🥺,,👉👈💖💖✨✨🥺,,,,👉👈💖💖💖✨✨🥺,👉👈🫂✨✨🥺,,👉👈💖💖✨✨✨👉👈💖💖✨✨✨✨👉👈"},
}

func main(t *testing.T) {
	for _, test := range testCases {
		encoded := ""

		encoded = Encode(test.value)

		if encoded != test.expected {
			t.Fatalf("%s failed to encode into %s. Found %s instead.", test.value, test.expected, encoded)
		} else {
			decoded, err := Decode(test.expected)
			if err != nil {
				t.Fatalf("%s failed to decode into %s. Found %s instead.", test.expected, test.value, decoded)
			}
		}
	}
}
