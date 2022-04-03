package main

import "testing"

func TestHello(t *testing.T) {

	// *testing.T と *testing.B の両方が満たすインターフェースである testing.TB を受け入れる
	assertCorrectMessage := func(t testing.TB, got, want string) {

		// メソッドがヘルパーであることをテストスイートに伝える
		// テストが失敗したときに報告される行番号は、テストヘルパーの中ではなく 呼び出された関数 の中を示す
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Test")
		want := "Hello, Test"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
