package main

import (
	"bytes"
	"testing"
)

func TestDeterministicOutputAcrossRepeatedRuns(t *testing.T) {
	const runs = 5
	input := []string{"testdata/valid_three_pieces_complex.txt"}
	expected := ""

	for i := 0; i < runs; i++ {
		var buffer bytes.Buffer
		executeWithWriter(input, &buffer)

		if i == 0 {
			expected = buffer.String()
			continue
		}

		if got := buffer.String(); got != expected {
			t.Fatalf("run %d produced different output: got %q want %q", i+1, got, expected)
		}
	}
}

func TestRegressionRejectsExtraBlankSeparator(t *testing.T) {
	var buffer bytes.Buffer

	executeWithWriter([]string{"testdata/invalid_format_extra_blank.txt"}, &buffer)

	if got := buffer.String(); got != "ERROR\n" {
		t.Fatalf("unexpected extra-blank regression output: got %q", got)
	}
}

func TestRegressionRejectsDisconnectedShape(t *testing.T) {
	var buffer bytes.Buffer

	executeWithWriter([]string{"testdata/invalid_shape_disconnected.txt"}, &buffer)

	if got := buffer.String(); got != "ERROR\n" {
		t.Fatalf("unexpected disconnected-shape regression output: got %q", got)
	}
}
