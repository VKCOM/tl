package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)

func SprintHexDump(data []byte) string {
	var buf bytes.Buffer
	buf.Grow(len(data) + len(data)/4)
	for i := 0; i < len(data); i += 4 {
		// Печатаем октеты в обратном порядке, чтобы они совпадали
		// с константами из `constant.go`.
		_, _ = fmt.Fprintf(&buf, "%02x%02x%02x%02x ",
			data[i+3],
			data[i+2],
			data[i+1],
			data[i+0])
	}
	return strings.TrimSpace(buf.String())
}

func SprintHexDumpTL2(data []byte) string {
	var buf bytes.Buffer
	buf.Grow(len(data) + len(data))
	for i := 0; i < len(data); i += 1 {
		// Печатаем октеты в обратном порядке, чтобы они совпадали
		// с константами из `constant.go`.
		_, _ = fmt.Fprintf(&buf, "%02x ",
			data[i+0])
	}
	return strings.TrimSpace(buf.String())
}

func ParseHexToBytes(data string) []byte {
	var result []byte
	for _, octet := range strings.Split(data, " ") {
		b1, _ := hex.DecodeString(octet[6:8])
		b2, _ := hex.DecodeString(octet[4:6])
		b3, _ := hex.DecodeString(octet[2:4])
		b4, _ := hex.DecodeString(octet[0:2])
		result = append(result, b1...)
		result = append(result, b2...)
		result = append(result, b3...)
		result = append(result, b4...)
	}
	return result
}

func ParseHexToBytesTL2(data string) []byte {
	var result []byte
	for _, octet := range strings.Split(data, " ") {
		b1, _ := hex.DecodeString(octet)
		result = append(result, b1...)
	}
	return result
}
