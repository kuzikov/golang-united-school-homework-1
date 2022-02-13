package solution

import "testing"

func TestGetMessage(t *testing.T) {
	want := "Hello 🗺️ !"
	if GetMessage() != want {
		t.Error()
	}
}
