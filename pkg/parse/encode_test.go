package parse

import "testing"

type testEncoderArgs struct {
	secret      string
	result      string
	expectError bool
}

var encoderTests = []testEncoderArgs{
	{
		secret:      "topsecret",
		result:      "dG9wc2VjcmV0",
		expectError: false,
	},
	{
		secret:      "extremelysecret",
		result:      "ZXh0cmVtZWx5c2VjcmV0",
		expectError: false,
	},
	{
		secret:      "uncrackable",
		result:      "dW5jcmFja2FibGU=",
		expectError: false,
	},
	{
		secret:      "iamwrong",
		result:      "Zmphc2xka2ZqYWpmc2xka2dqYWxzZGtmanNkZmE=",
		expectError: true,
	},
}

func TestEncode(t *testing.T) {
	p := NewParser()

	for _, test := range encoderTests {
		if res, err := p.Encode(test.secret); err != nil && !test.expectError || res != test.result && !test.expectError {
			t.Errorf("Result %q is not equal to expected %q", res, test.result)
		}
	}
}
