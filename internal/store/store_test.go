package store

import "testing"

func TestStore(t *testing.T) {
	s := New()
	var testGetHello = func(want string) {
		if got := s.Get("hello"); got != want {
			t.Errorf("got \"%s\", want \"%s\"\n", got, want)
		}
	}

	testGetHello("")

	s.Add("hello", "world")
	testGetHello("world")

	s.Delete("hello")
	testGetHello("")
}
