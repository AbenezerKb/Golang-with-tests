package code

import (
	"bytes"
	"testing"
)

func TestDependancy(t *testing.T) {

	responsePrinter := bytes.Buffer{}
	Printer(&responsePrinter, "Printer passing")

	got := responsePrinter.String()
	want := "Printer passing, Printer responding"

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}

}
