package sandbox

import "testing"

func TestFoo(t *testing.T) {
	want := "Hello, Gopher!"
	if got := Foo("Gopher"); got != want {
		t.Fatalf("Foo() = %q, want %q", got, want)
	}
}

func TestBar(t *testing.T) {
	if Bar(2, 3) != 5 {
		t.Fatalf("Bar() failed")
	}
}
