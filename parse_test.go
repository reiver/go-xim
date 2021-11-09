package xim

import (
	"testing"
)

func TestParse(t *testing.T) {

	for testNumber, test := range stdtests {

		var id ID = something(test.Value)

		var ximid string = id.String()

		var actual ID = Parse(ximid)
		if Nothing() == actual {
			t.Errorf("For test #%d, the actual result of parsing is not what was expected — it is option-type value nothing" , testNumber)
			t.Logf("VALUE:    %#064b", test.Value)
			t.Logf("ID:     %s", id)
			{
				chaos, successful := id.Chaos()
				t.Logf("ID-CHAOS-SUCCESSFUL: %t", successful)
				t.Logf("ID-CHAOS: %#064b", chaos)
			}
			t.Logf("XIM-ID: %s", ximid)
			t.Logf("ACTUAL: %s", actual)
			continue
		}

		{
			actualchaos, successful := id.Chaos()
			if !successful {
				t.Errorf("For test #%d, the actual result of parsing is not what was expected — it is option-type value nothing" , testNumber)
				t.Logf("VALUE:    %#064b", test.Value)
				t.Logf("ID:     %s", id)
				{
					chaos, successful := id.Chaos()
					t.Logf("ID-CHAOS-SUCCESSFUL: %t", successful)
					t.Logf("ID-CHAOS: %#064b", chaos)
				}
				t.Logf("XIM-ID: %s", ximid)
				t.Logf("ACTUAL: %s", actual)
				continue
			}

			if expectedchaos := test.Value & maskchaos; expectedchaos != actualchaos {
				t.Errorf("For test #%d, the chaos from the actual result of parsing is not what was expected" , testNumber)
				t.Logf("VALUE:          %#064b", test.Value)
				t.Logf("ID:     %s", id)
				t.Logf("XIM-ID: %s", ximid)
				t.Logf("ACTUAL: %s", actual)
				t.Logf("ACTUAL-CHAOS:   %#064b", actualchaos)
				t.Logf("EXPECTED-CHAOS: %#064b", expectedchaos)
				continue
			}
		}
	}
}
