package tl2pure

import "github.com/vkcom/tl/pkg/basictl"

type ByteColor byte

const ByteColorNormal = 0
const ByteColorObjectSize = 1
const ByteColorElementCount = 2
const ByteColorVariantIndex = 3
const ByteColorFieldMask = 4

type ByteBuilder struct {
	buf    []byte
	colors []ByteColor
}

func (b *ByteBuilder) Buf() []byte {
	return b.buf
}

func (b *ByteBuilder) Len() int {
	return len(b.buf)
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
	b.fixColors(ByteColorObjectSize)
	b.buf = append(b.buf, v...)
	b.fixColors(ByteColorNormal)
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
	offset += copy(b.buf[firstUsedByte-16+offset:], b.buf[firstUsedByte:lastUsedByte])
	copy(b.colors[firstUsedByte-16+offset:], b.colors[firstUsedByte:lastUsedByte])
	b.buf = b.buf[:firstUsedByte-16+offset]
	b.colors = b.colors[:firstUsedByte-16+offset]
}

// at current offset
func (b *ByteBuilder) SetCursor() {

}

func (b *ByteBuilder) PushStyle(s UIStyle) {

}

func (b *ByteBuilder) PopStyle() {

}
