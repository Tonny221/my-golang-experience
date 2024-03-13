package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Tonny")
  want := "Hello, Tonny"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
