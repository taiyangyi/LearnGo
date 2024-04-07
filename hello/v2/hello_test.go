package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello world"

	if got != want {
		t.Error("测试失败...")
		t.Errorf("got %q want %q", got, want)
	}
}

// PS F:\go\goprojects\LearnGo\hello\v2> go test
// PASS
// ok      learngo/hello/v2        0.480s

// PS F:\go\goprojects\LearnGo\hello\v2> go test
// --- FAIL: TestHello (0.00s)
//     hello_test.go:10: 测试失败...
//     hello_test.go:11: got "hello, world" want "hello world"
// FAIL
// exit status 1
// FAIL    learngo/hello/v2        0.563s
