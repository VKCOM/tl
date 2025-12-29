package tl2pure

import "github.com/vkcom/tl/pkg/basictl"

type ByteBuilder struct {
	buf []byte
}

func (b *ByteBuilder) Buf() []byte {
	return b.buf
}

func (b *ByteBuilder) Len() int {
	return len(b.buf)
}

func (b *ByteBuilder) Append(w ...byte) {
	b.buf = append(b.buf, w...)
}

func (b *ByteBuilder) WriteString(v string) {
	b.buf = basictl.StringWriteTL2(b.buf, v)
}

func (b *ByteBuilder) WriteVariantIndex(v int) {
	b.buf = basictl.TL2WriteSize(b.buf, v)
}

func (b *ByteBuilder) WriteElementCount(v int) {
	b.buf = basictl.TL2WriteSize(b.buf, v)
}

func (b *ByteBuilder) ReserveSpaceForSize() int {
	b.buf = append(b.buf, make([]byte, 16)...)
	return len(b.buf)
}

func (b *ByteBuilder) FinishSize(firstUsedByte int, lastUsedByte int, optimizeEmpty bool) {
	if optimizeEmpty && firstUsedByte == lastUsedByte {
		b.buf = b.buf[:firstUsedByte-16]
		return
	}
	offset := basictl.TL2PutSize(b.buf[firstUsedByte-16:], lastUsedByte-firstUsedByte)
	offset += copy(b.buf[firstUsedByte-16+offset:], b.buf[firstUsedByte:lastUsedByte])
	b.buf = b.buf[:firstUsedByte-16+offset]
}

// at current offset
func (b *ByteBuilder) SetCursor() {

}

func (b *ByteBuilder) PushStyle(s UIStyle) {

}

func (b *ByteBuilder) PopStyle() {

}
