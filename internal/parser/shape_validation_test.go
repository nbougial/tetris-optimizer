package parser

import "testing"

func TestValidateBlockHashCounts(t *testing.T) {
	validBlocks := [][]string{{"##..", "##..", "....", "...."}}
	if err := ValidateBlockHashCounts(validBlocks); err != nil {
		t.Fatalf("expected valid hash count, got error: %v", err)
	}

	invalidBlocks := [][]string{{"###.", "##..", "....", "...."}}
	if err := ValidateBlockHashCounts(invalidBlocks); err == nil {
		t.Fatal("expected invalid hash count error")
	}
}

func TestValidateBlockConnectivityAcceptsConnectedShape(t *testing.T) {
	blocks := [][]string{{"##..", ".##.", "....", "...."}}

	if err := ValidateBlockConnectivity(blocks); err != nil {
		t.Fatalf("expected connected shape to pass, got error: %v", err)
	}
}

func TestValidateBlockConnectivityRejectsDisconnectedShape(t *testing.T) {
	blocks := [][]string{{"#.#.", ".#..", ".#..", "...."}}

	if err := ValidateBlockConnectivity(blocks); err == nil {
		t.Fatal("expected disconnected shape error")
	}
}
