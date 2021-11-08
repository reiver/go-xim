package xim

import (
	"testing"
)

func TestDecompile(t *testing.T) {

	for testNumber, test := range stdtests {

		first, second := decompile(test.Value)

		actual := compile(first, second)

		if expected := test.Value; expected != actual {
			t.Errorf("For test #%d, the actual re-compiled value of the decompiled values is not what was expected.", testNumber)
			t.Logf("EXPECTED: %064b", expected)
			t.Logf("ACTUAL:   %064b", actual)
			t.Logf("FIRST     %064b", first)
			t.Logf("SECOND    %064b", second)
			continue
		}
	}
}
