package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"strings"
)

type IntelHexSummary struct {
	Records              int
	DataRecords          int
	DataBytes            int
	HasEOF               bool
	ExtendedAddressBase  uint32
}

// ParseIntelHex validates Intel HEX record lengths and checksums.
func ParseIntelHex(input string) (IntelHexSummary, error) {
	var summary IntelHexSummary
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if summary.HasEOF {
			return summary, fmt.Errorf("line %d: data found after EOF record", lineNumber)
		}
		if !strings.HasPrefix(line, ":") {
			return summary, fmt.Errorf("line %d: record must start with ':'", lineNumber)
		}

		record, err := hex.DecodeString(line[1:])
		if err != nil {
			return summary, fmt.Errorf("line %d: invalid hexadecimal data: %w", lineNumber, err)
		}
		if len(record) < 5 {
			return summary, fmt.Errorf("line %d: record is shorter than the minimum length", lineNumber)
		}

		byteCount := int(record[0])
		if len(record) != byteCount+5 {
			return summary, fmt.Errorf("line %d: byte count is %d but record contains %d data bytes", lineNumber, byteCount, len(record)-5)
		}

		var checksum byte
		for _, value := range record {
			checksum += value
		}
		if checksum != 0 {
			return summary, fmt.Errorf("line %d: checksum mismatch", lineNumber)
		}

		address := uint16(record[1])<<8 | uint16(record[2])
		recordType := record[3]
		data := record[4 : 4+byteCount]

		switch recordType {
		case 0x00: // Data record
			summary.DataRecords++
			summary.DataBytes += byteCount
		case 0x01: // End-of-file record
			if byteCount != 0 || address != 0 {
				return summary, fmt.Errorf("line %d: invalid EOF record", lineNumber)
			}
			summary.HasEOF = true
		case 0x02: // Extended segment address
			if byteCount != 2 {
				return summary, fmt.Errorf("line %d: extended segment address must contain 2 bytes", lineNumber)
			}
			summary.ExtendedAddressBase = uint32(data[0])<<12 | uint32(data[1])<<4
		case 0x04: // Extended linear address
			if byteCount != 2 {
				return summary, fmt.Errorf("line %d: extended linear address must contain 2 bytes", lineNumber)
			}
			summary.ExtendedAddressBase = (uint32(data[0])<<8 | uint32(data[1])) << 16
		case 0x03, 0x05: // Start segment / linear address
			if byteCount != 4 {
				return summary, fmt.Errorf("line %d: start address record must contain 4 bytes", lineNumber)
			}
		default:
			return summary, fmt.Errorf("line %d: unsupported record type 0x%02X", lineNumber, recordType)
		}

		summary.Records++
	}

	if err := scanner.Err(); err != nil {
		return summary, fmt.Errorf("read Intel HEX data: %w", err)
	}
	if summary.Records == 0 {
		return summary, fmt.Errorf("file contains no Intel HEX records")
	}
	if !summary.HasEOF {
		return summary, fmt.Errorf("file is missing the end-of-file record")
	}

	return summary, nil
}
