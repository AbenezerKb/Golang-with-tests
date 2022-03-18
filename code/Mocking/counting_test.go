package code

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {

	buffer := &bytes.Buffer{}

	countdown(buffer)

	want := "3"
	got := buffer.String()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
