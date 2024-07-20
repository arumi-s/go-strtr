package strtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrtr(t *testing.T) {
	// should replace substrings correctly
	assert.Equal(t, "X Planet", Strtr("Hello World", map[string]string{"Hello": "X", "World": "Planet"}))

	// should not search replaced values again
	assert.Equal(t, "xxxX World", Strtr("xxxHello World", map[string]string{"Hello": "X", "X": "Something else"}))

	// should handle overlapping keys with longest match being replaced first
	assert.Equal(t, "a234", Strtr("abcde", map[string]string{"abc": "123", "bcde": "234"}))

	// should manage multiple replacements without overlaps
	assert.Equal(t, "Hi Earth", Strtr("Hello World", map[string]string{"Hello": "Hi", "World": "Earth"}))

	// should work when no keys are found
	assert.Equal(t, "Hello World", Strtr("Hello World", map[string]string{"Test": "X", "Dummy": "Y"}))

	// should prioritize longer substrings in replacements
	assert.Equal(t, "xy", Strtr("ababc", map[string]string{"ab": "x", "abc": "y"}))
	assert.Equal(t, "by", Strtr("babc", map[string]string{"ab": "x", "abc": "y"}))

	// should correctly replace when substrings are prefixes of each other
	assert.Equal(t, "123ef", Strtr("abcdef", map[string]string{"ab": "AB", "abc": "XYZ", "abcd": "123"}))

	// should handle cases where replaced substrings could match other keys
	assert.Equal(t, "123ana", Strtr("banana", map[string]string{"ban": "123", "123": "xyz"}))

	// should handle special characters and spaces
	assert.Equal(t, "Hi Earth.", Strtr("Hello, World!", map[string]string{"Hello": "Hi", ", World!": " Earth."}))

	// should handle unicode characters in replacements
	assert.Equal(t, "HelloWorld", Strtr("こんにちは世界", map[string]string{"こんにちは": "Hello", "世界": "World"}))

	// should handle complex nested replacements correctly
	assert.Equal(t, "begin center stop", Strtr("start middle end", map[string]string{
		"start":             "begin",
		"middle":            "center",
		"end":               "stop",
		"begin center stop": "complete",
	}))

	// should handle large length changes
	assert.Equal(t, "defghijklmnghi123456789045i", Strtr("abcdefghi", map[string]string{
		"abc": "defghijklmnghi",
		"def": "1234567890",
		"gh":  "45",
	}))
	assert.Equal(t, "a slow brown dog jumps over fox", Strtr("a quick brown fox jumps over the lazy dog", map[string]string{
		"the lazy dog": "fox",
		"fox":          "dog",
		"quick":        "slow",
	}))
}
