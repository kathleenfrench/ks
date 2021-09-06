package parse

import "testing"

type keyMapTest struct {
	input         map[string]string
	result        []string
	expectNoMatch bool
}

var keyMapTests = []keyMapTest{
	{
		input: map[string]string{
			"test":    "hello",
			"world":   "hi",
			"another": "command line tool",
		},
		result: []string{
			"test,", "world", "another",
		},
		expectNoMatch: true,
	},
	{
		input: map[string]string{
			"here is a test":      "and a value",
			"hello":               "again",
			"testing testing 123": "correct!",
		},
		result: []string{
			"here is a test",
			"hello",
			"testing testing 123",
		},
		expectNoMatch: false,
	},
}

func TestGetMapKeys(t *testing.T) {
	p := NewParser()

	for _, test := range keyMapTests {
		res := p.GetMapKeys(test.input)
		match := equal(res, test.result)
		if !match && !test.expectNoMatch {
			t.Errorf("Result %q is not equal to expected %q", res, test.result)
		}
	}
}

func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	diff := make(map[string]int, len(a))

	for _, aa := range a {
		diff[aa]++
	}

	for _, bb := range b {
		if _, ok := diff[bb]; !ok {
			return false
		}

		diff[bb]--
		if diff[bb] == 0 {
			delete(diff, bb)
		}
	}

	return len(diff) == 0
}
