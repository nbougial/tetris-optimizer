package main

import (
	"bytes"
	"testing"
)

func TestWriteErrorOutputsExactMessageWithNewline(t *testing.T) {
	var buffer bytes.Buffer

	writeError(&buffer)

	if got := buffer.String(); got != "ERROR\n" {
		t.Fatalf("unexpected error output: got %q want %q", got, "ERROR\n")
	}
}

func TestExecuteWithWriterPrintsOnlyErrorForInvalidArgs(t *testing.T) {
	var buffer bytes.Buffer

	executeWithWriter([]string{}, &buffer)

	if got := buffer.String(); got != "ERROR\n" {
		t.Fatalf("unexpected execute output: got %q want %q", got, "ERROR\n")
	}
}

func TestExecuteWithWriterStaysSilentOnSuccessfulRun(t *testing.T) {
	var buffer bytes.Buffer

	executeWithWriter([]string{"testdata/task12_valid_two_pieces.txt"}, &buffer)

	if got := buffer.String(); got != "" {
		t.Fatalf("expected no output on success, got %q", got)
	}
}
