package byteconst

import (
	"testing"
)

func TestConsts(t *testing.T) {
	if got := KB; got != 1000 {
		t.Errorf("KB should be 1000, but is %d", got)
	}
	if got := MB; got != 1000000 {
		t.Errorf("MB should be 1,000,000; but is %d", got)
	}
	if got := GB; got != 1000000000 {
		t.Errorf("KB should be 1,000,000,000; but is %d", got)
	}
}
