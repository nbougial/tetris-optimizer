package main

import (
	"bytes"
	"testing"
)

func TestCLIIntegrationRejectsNoArgs(t *testing.T) {
	var buffer bytes.Buffer

	executeWithWriter([]string{}, &buffer)

	if got := buffer.String(); got != "ERROR\n" {
		t.Fatalf("unexpected no-args output: got %q", got)
	}
}

func TestCLIIntegrationRejectsExtraArgs(t *testing.T) {
	var buffer bytes.Buffer

	executeWithWriter([]string{"a", "b"}, &buffer)

	if got := buffer.String(); got != "ERROR\n" {
		t.Fatalf("unexpected extra-args output: got %q", got)
	}
}

func TestCLIIntegrationRejectsMissingFile(t *testing.T) {
	var buffer bytes.Buffer

	executeWithWriter([]string{"testdata/does-not-exist.txt"}, &buffer)

	if got := buffer.String(); got != "ERROR\n" {
		t.Fatalf("unexpected missing-file output: got %q", got)
	}
}

func TestCLIIntegrationRendersSolvedBoard(t *testing.T) {
	var buffer bytes.Buffer

	executeWithWriter([]string{"testdata/task12_valid_two_pieces.txt"}, &buffer)

	want := "AAB.\nAAB.\n..B.\n..B.\n"
	if got := buffer.String(); got != want {
		t.Fatalf("unexpected solved board output: got %q want %q", got, want)
	}
}
