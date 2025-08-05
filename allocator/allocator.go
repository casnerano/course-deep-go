package allocator

import (
	"unsafe"
)

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	start := uintptr(unsafe.Pointer(&memory[0]))
	for ptrIndex, ptr := range pointers {
		offset := uintptr(ptr) - start
		if offset != uintptr(ptrIndex) {
			memory[ptrIndex], memory[offset] = memory[offset], 0
		}
		pointers[ptrIndex] = unsafe.Pointer(&memory[ptrIndex])
	}
}
