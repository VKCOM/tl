package tl2pure

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/pkg/basictl"
)

type ByteColor byte

const ByteColorNormal ByteColor = 0
const ByteColorObjectSize ByteColor = 1
const ByteColorElementCount ByteColor = 2
const ByteColorVariantIndex ByteColor = 3
const ByteColorFieldMask ByteColor = 4
const ByteColorStringSize ByteColor = 6
const ByteColorStringData ByteColor = 7

type ByteBuilder struct {
	buf          []byte
	colors       []ByteColor
	cursorStart  int
	cursorFinish int
}

func (b *ByteBuilder) Buf() []byte {
	return b.buf
}

func (b *ByteBuilder) Len() int {
	return len(b.buf)
}

func (b *ByteBuilder) PrintLegend() string {
	var sb strings.Builder
	sb.WriteString("Legend: ")
	sb.WriteString(color.Green)
	sb.WriteString("object size ")
	sb.WriteString(color.Blue)
	sb.WriteString("array len ")
	sb.WriteString(color.Red)
	sb.WriteString("union variant ")
	sb.WriteString(color.Purple)
	sb.WriteString("fieldmask ")
	sb.WriteString(color.Yellow)
	sb.WriteString("string len ")
	sb.WriteString(color.Gray)
	sb.WriteString("string data")
	sb.WriteString(color.Reset)
	return sb.String()
}

func (b *ByteBuilder) Print() string {
	var sb strings.Builder
	currentColor := ByteColorNormal
	currentCursor := false
	for i, bt := range b.buf {
		wantColor := b.colors[i]
		wantCursor := i >= b.cursorStart && i < b.cursorFinish
		if !wantCursor && currentCursor {
			sb.WriteString(color.Reset)
			currentColor = ByteColorNormal
			currentCursor = false
		}
		if wantColor == ByteColorNormal && currentColor != ByteColorNormal {
			sb.WriteString(color.Reset)
			currentColor = ByteColorNormal
			currentCursor = false
		}
		if wantCursor && !currentCursor {
			sb.WriteString(color.Underline)
			currentCursor = true
		}
		if wantColor != currentColor {
			switch wantColor {
			case ByteColorObjectSize:
				sb.WriteString(color.Green)
			case ByteColorElementCount:
				sb.WriteString(color.Blue)
			case ByteColorVariantIndex:
				sb.WriteString(color.Red)
			case ByteColorFieldMask:
				sb.WriteString(color.Purple)
			case ByteColorStringSize:
				sb.WriteString(color.Yellow)
			case ByteColorStringData:
				sb.WriteString(color.Gray)
			}
			currentColor = wantColor
		}
		_, _ = fmt.Fprintf(&sb, "%02x ", bt)
	}
	if currentColor != ByteColorNormal || currentCursor {
		sb.WriteString(color.Reset)
	}
	return sb.String()
}

// primitives write directly into buf, so we call it at the start of each our fun
func (b *ByteBuilder) fixColors(c ByteColor) {
	if len(b.colors) > len(b.buf) {
		panic("color markers must never be ahead of bytes")
	}
	for len(b.colors) < len(b.buf) {
		b.colors = append(b.colors, c)
	}
}

func (b *ByteBuilder) WriteFieldmask() {
	b.fixColors(ByteColorNormal)
	b.buf = append(b.buf, 0)
	b.fixColors(ByteColorFieldMask)
}

func (b *ByteBuilder) WriteString(v string) {
	b.fixColors(ByteColorNormal)
	b.buf = basictl.TL2WriteSize(b.buf, len(v))
	b.fixColors(ByteColorStringSize)
	b.buf = append(b.buf, v...)
	b.fixColors(ByteColorStringData)
}

func (b *ByteBuilder) WriteVariantIndex(v int) {
	b.fixColors(ByteColorNormal)
	b.buf = basictl.TL2WriteSize(b.buf, v)
	b.fixColors(ByteColorVariantIndex)
}

func (b *ByteBuilder) WriteElementCount(v int) {
	b.fixColors(ByteColorNormal)
	b.buf = basictl.TL2WriteSize(b.buf, v)
	b.fixColors(ByteColorElementCount)
}

func (b *ByteBuilder) ReserveSpaceForSize() int {
	b.fixColors(ByteColorNormal)
	b.buf = append(b.buf, make([]byte, 16)...)
	b.fixColors(ByteColorObjectSize) // mark all bytes as size for simplicity of FinishSize
	return len(b.buf)
}

func (b *ByteBuilder) FinishSize(firstUsedByte int, lastUsedByte int, optimizeEmpty bool) {
	b.fixColors(ByteColorNormal)
	if optimizeEmpty && firstUsedByte == lastUsedByte {
		b.buf = b.buf[:firstUsedByte-16]
		b.colors = b.colors[:firstUsedByte-16]
		return
	}
	offset := basictl.TL2PutSize(b.buf[firstUsedByte-16:], lastUsedByte-firstUsedByte)
	if b.cursorStart > firstUsedByte {
		b.cursorStart -= 16 - offset
	}
	if b.cursorFinish > firstUsedByte {
		b.cursorFinish -= 16 - offset
	}
	copy(b.colors[firstUsedByte-16+offset:], b.colors[firstUsedByte:lastUsedByte])
	offset += copy(b.buf[firstUsedByte-16+offset:], b.buf[firstUsedByte:lastUsedByte])
	b.buf = b.buf[:firstUsedByte-16+offset]
	b.colors = b.colors[:firstUsedByte-16+offset]
}

func (b *ByteBuilder) SetCursorStart() {
	b.cursorStart = len(b.buf)
}

func (b *ByteBuilder) SetCursorFinish() {
	b.cursorFinish = len(b.buf)
}
