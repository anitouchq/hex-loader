package main

import "testing"

const validIntelHex = ":0400000001020304F2\n:00000001FF\n"

func TestParseIntelHexCountsDataBytes(t *testing.T) {
	got, err := ParseIntelHex(validIntelHex)
	if err != nil {
		t.Fatalf("ParseIntelHex() error = %v", err)
	}
	if got.DataBytes != 4 {
		t.Fatalf("DataBytes = %d, want 4", got.DataBytes)
	}
	if got.DataRecords != 1 {
		t.Fatalf("DataRecords = %d, want 1", got.DataRecords)
	}
	if !got.HasEOF {
		t.Fatal("HasEOF = false, want true")
	}
}

func TestParseIntelHexRejectsInvalidChecksum(t *testing.T) {
	input := ":0400000001020304F3\n:00000001FF\n"
	if _, err := ParseIntelHex(input); err == nil {
		t.Fatal("ParseIntelHex() accepted invalid checksum")
	}
}

func TestParseIntelHexRequiresEOFRecord(t *testing.T) {
	input := ":0400000001020304F2\n"
	if _, err := ParseIntelHex(input); err == nil {
		t.Fatal("ParseIntelHex() accepted file without EOF record")
	}
}

func TestParseIntelHexRejectsDataAfterEOF(t *testing.T) {
	input := ":00000001FF\n:0400000001020304F2\n"
	if _, err := ParseIntelHex(input); err == nil {
		t.Fatal("ParseIntelHex() accepted data after EOF")
	}
}
