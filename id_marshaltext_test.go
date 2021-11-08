package xim

import (
	"testing"
)

func TestID_MarshalText(t *testing.T) {

	for testNumber, test := range stdtests {

		var id ID = something(test.Value)

		var marshaled []byte
		{
			var err error

			marshaled, err = id.MarshalText()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error when mashaling, but actually got one.", testNumber)
				t.Logf("VALUE: %064b", test.Value)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
			if nil == marshaled {
				t.Errorf("For test #%d, the actual marshaled value is nil but that was not expected.", testNumber)
				t.Logf("VALUE: %064b", test.Value)
				t.Logf("MARSHALED: %s", marshaled)
				continue
			}
		}

		{
			var newid ID

			err := newid.UnmarshalText(marshaled)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error when unmashaling, but actually got one.", testNumber)
				t.Logf("VALUE: %064b", test.Value)
				t.Logf("MARSHALED: %s", marshaled)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}


			var expected ID = id
			var actual   ID = newid

			if expected != actual {
				t.Errorf("For test #%d, the actual unmarshaled marshaled value is not what was expected.", testNumber)
				t.Logf("VALUE: %064b", test.Value)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				continue
			}

		}
	}
}
