package lb

import "testing"

func TestBoxId(t *testing.T) {
	if b := GetBox(7, 7); b != 8 {
		t.Error()
	}
}
