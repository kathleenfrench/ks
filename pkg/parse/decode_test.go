package parse

import "testing"

type testDecoderArgs struct {
	secret      string
	result      string
	expectError bool
}

var decoderTests = []testDecoderArgs{
	{
		secret:      "dG9wc2VjcmV0",
		result:      "topsecret",
		expectError: false,
	},
	{
		secret:      "ZXh0cmVtZWx5c2VjcmV0",
		result:      "extremelysecret",
		expectError: false,
	},
	{
		secret:      "dW5jcmFja2FibGU=",
		result:      "iamwrong",
		expectError: true,
	},
}

func TestDecode(t *testing.T) {
	p := NewParser()

	for _, test := range decoderTests {
		if res, err := p.Decode(test.secret); err != nil && !test.expectError || res != test.result && !test.expectError {
			t.Errorf("Result %q is not equal to expected %q", res, test.result)
		}
	}
}
