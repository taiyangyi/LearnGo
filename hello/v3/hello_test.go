package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("HangZhou")
	want := "hello HangZhou"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

// PS F:\go\goprojects\LearnGo\hello\v3> go test
// PASS
// ok      learngo/hello/v3        1.517s
