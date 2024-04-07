package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Go")
	want := "Hello,Go"

	if got != want {
		t.Error("测试失败...")
		t.Errorf("got %q want %q", got, want)
	}
}
