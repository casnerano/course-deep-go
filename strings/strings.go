package strings

import (
	"runtime"
	"unsafe"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func newCOWBuffer(data []byte, refs *int) *COWBuffer {
	if refs == nil {
		defRefCount := 1
		refs = &defRefCount
	}

	buf := COWBuffer{
		data: data,
		refs: refs,
	}

	runtime.SetFinalizer(&buf, (*COWBuffer).Close)
	return &buf
}

func NewCOWBuffer(data []byte) COWBuffer {
	return *newCOWBuffer(data, nil)
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++
	return *newCOWBuffer(b.data, b.refs)
}

func (b *COWBuffer) Close() {
	if b.refs != nil && *b.refs > 1 {
		*b.refs--

		b.refs = nil
		b.data = nil

		runtime.SetFinalizer(b, nil)
	}
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}

	if *b.refs > 1 {
		t := make([]byte, len(b.data))
		copy(t, b.data)
		b.data = t
		*b.refs--
	}

	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
