package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Perdi", "")
		want := "Hello, Perdi"
	
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello world' when an empty string supplied", func(t *testing.T) {
		got:= Hello("",  "")
		want:= "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Indonesia", func(t *testing.T) {
		got:= Hello("Perdi", "Indonesia")
		want := "Hi, Perdi"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got!= want {
		t.Errorf("Got %q want %q", got, want)
	}

}