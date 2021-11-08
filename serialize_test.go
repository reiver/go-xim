package xim

import (
	"testing"
)

func TestSerialize(t *testing.T) {

	for testNumber, test := range stdtests {

		var serialized string = serialize(test.Value)

		actual, successful := unserialize(serialized)
		if !successful {
			t.Errorf("For test #%d, expected unserialization of serialized data to be successful but wasn't." , testNumber)
			t.Logf("SERIALIZED: %s", serialized)
			t.Logf("SUCCESSFUL: %t", successful)
			t.Logf("VALUE: %064b", test.Value)
			continue
		}


		if expected := test.Value; expected != actual {
			t.Errorf("For test #%d, ", testNumber)
			t.Logf("SERIALIZED: %s", serialized)
			t.Logf("SUCCESSFUL: %t", successful)
			t.Logf("EXPECTED: %064b", expected)
			t.Logf("ACTUAL:   %064b", actual)
			continue
		}
	}
}
