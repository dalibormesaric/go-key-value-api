package history

import (
	"testing"
)

func TestHistory(t *testing.T) {
	h := New()
	var testGetHello = func(want string, wantAmount int, wantState string) {
		r := h.Get("hello")
		if len(r) != wantAmount {
			t.Errorf("got %d item(s), want %d\n", len(r), wantAmount)
		}
		if len(r) > 0 && r[0].Value != want {
			t.Errorf("got value \"%s\", want \"%s\"\n", r[0].Value, want)
		}
		if len(r) > 0 && r[0].State != wantState {
			t.Errorf("got state = \"%s\", want \"%s\"\n", r[0].State, wantState)
		}
	}

	testGetHello("", 0, "")

	h.Append("hello", "world")
	testGetHello("world", 1, added)

	h.Append("hello", "again")
	testGetHello("again", 2, updated)

	h.Delete("hello")
	testGetHello("", 3, deleted)

	h.Append("hello", "there")
	testGetHello("there", 4, added)

	h.Append("hello", "you")
	testGetHello("you", 5, updated)
}
