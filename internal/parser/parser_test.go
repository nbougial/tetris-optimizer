package parser

import "testing"

func TestSplitRawBlocksSingleBlock(t *testing.T) {
	content := []byte("##..\n##..\n....\n....\n")

	blocks, err := SplitRawBlocks(content)
	if err != nil {
		t.Fatalf("expected single block to parse, got error: %v", err)
	}

	if len(blocks) != 1 {
		t.Fatalf("expected 1 block, got %d", len(blocks))
	}

	if len(blocks[0]) != 4 {
		t.Fatalf("expected block to contain 4 lines, got %d", len(blocks[0]))
	}
}

func TestSplitRawBlocksMultipleBlocksPreservesOrder(t *testing.T) {
	content := []byte("##..\n##..\n....\n....\n\n#...\n#...\n#...\n#...\n")

	blocks, err := SplitRawBlocks(content)
	if err != nil {
		t.Fatalf("expected multiple blocks to parse, got error: %v", err)
	}

	if len(blocks) != 2 {
		t.Fatalf("expected 2 blocks, got %d", len(blocks))
	}

	if blocks[0][0] != "##.." || blocks[1][0] != "#..." {
		t.Fatalf("expected blocks to preserve input order, got %#v", blocks)
	}
}

func TestSplitRawBlocksRejectsMalformedSeparators(t *testing.T) {
	testCases := []struct {
		name    string
		content []byte
	}{
		{
			name:    "leading blank line",
			content: []byte("\n##..\n##..\n....\n....\n"),
		},
		{
			name:    "consecutive blank lines",
			content: []byte("##..\n##..\n....\n....\n\n\n#...\n#...\n#...\n#...\n"),
		},
		{
			name:    "empty file",
			content: []byte(""),
		},
	}

	for _, testCase := range testCases {
		if _, err := SplitRawBlocks(testCase.content); err == nil {
			t.Fatalf("%s: expected malformed separators error", testCase.name)
		}
	}
}

func TestValidateBlockLineCounts(t *testing.T) {
	validBlocks := [][]string{{"##..", "##..", "....", "...."}}
	if err := ValidateBlockLineCounts(validBlocks); err != nil {
		t.Fatalf("expected valid line counts, got error: %v", err)
	}

	invalidBlocks := [][]string{{"##..", "##..", "...."}}
	if err := ValidateBlockLineCounts(invalidBlocks); err == nil {
		t.Fatal("expected invalid line count error")
	}
}

func TestValidateBlockColumnCounts(t *testing.T) {
	validBlocks := [][]string{{"##..", "##..", "....", "...."}}
	if err := ValidateBlockColumnCounts(validBlocks); err != nil {
		t.Fatalf("expected valid column counts, got error: %v", err)
	}

	invalidBlocks := [][]string{{"##..", "##...", "....", "...."}}
	if err := ValidateBlockColumnCounts(invalidBlocks); err == nil {
		t.Fatal("expected invalid column count error")
	}
}

func TestValidateBlockCharacters(t *testing.T) {
	validBlocks := [][]string{{"##..", "##..", "....", "...."}}
	if err := ValidateBlockCharacters(validBlocks); err != nil {
		t.Fatalf("expected valid characters, got error: %v", err)
	}

	invalidBlocks := [][]string{{"##..", "#x..", "....", "...."}}
	if err := ValidateBlockCharacters(invalidBlocks); err == nil {
		t.Fatal("expected invalid character error")
	}
}
