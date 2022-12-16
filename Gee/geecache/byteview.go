package geecache

type ByteView struct {
	b []byte
}

func newByte() *ByteView {
	return &ByteView{b: make([]byte, 0)}
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func (v ByteView) CloneBytes() []byte {
	bytes := make([]byte, len(v.b))
	copy(bytes, v.b)
	return bytes
}
