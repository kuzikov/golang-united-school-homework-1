package solution

import "testing"

func TestGetMessage(t *testing.T) {
	if GetMessage() != "Hello 🗺️!" {
		t.Error()
	}
}
