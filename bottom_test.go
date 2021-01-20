package bottom

import (
	"testing"
)

type testCase struct {
	value    string
	expected string
	bytes    bool
}

var testCases [4]testCase = [4]testCase{
	{"Test", "💖✨✨✨,,,,👉👈💖💖,👉👈💖💖✨🥺👉👈💖💖✨🥺,👉👈", false},
	{"h", "💖💖,,,,👉👈", true},
	{"🥺", "🫂✨✨✨✨👉👈💖💖💖🥺,,,,👉👈💖💖💖✨🥺👉👈💖💖💖✨✨✨🥺,👉👈", false},
	{"がんばれ", "🫂✨✨🥺,,👉👈💖💖✨✨🥺,,,,👉👈💖💖✨✨✨✨👉👈🫂✨✨🥺,,👉👈💖💖✨✨✨👉👈💖💖✨✨✨✨🥺,,👉👈🫂✨✨🥺,,👉👈💖💖✨✨🥺,,,,👉👈💖💖💖✨✨🥺,👉👈🫂✨✨🥺,,👉👈💖💖✨✨✨👉👈💖💖✨✨✨✨👉👈", false},
}

func main(t *testing.T) {
	for _, test := range testCases {
		encoded := ""

		if test.bytes {
			encoded = encodeBytes(test.value)
		} else {
			encoded = encode(test.value)
		}

		if encoded != test.expected {
			t.Fatalf("%s failed to encode into %s. Found %s instead.", test.value, test.expected, encoded)
		} else {
			decoded, err := decode(test.expected)
			if err != nil {
				t.Fatalf("%s failed to decode into %s. Found %s instead.", test.expected, test.value, decoded)
			}
		}
	}
}
