package garbage_collector

import (
	"unsafe"
)

const nullPtr = uintptr(0)

func Trace(stacks [][]uintptr) []uintptr {
	var trace []uintptr
	visited := make(map[uintptr]struct{})

	var traverse func(ptr uintptr)
	traverse = func(ptr uintptr) {
		if ptr == nullPtr {
			return
		}

		if _, ok := visited[ptr]; ok {
			return
		}

		trace = append(trace, ptr)
		visited[ptr] = struct{}{}

		nextPtr := *(*uintptr)(unsafe.Pointer(ptr))
		if nextPtr != nullPtr {
			traverse(nextPtr)
		}
	}

	for _, stack := range stacks {
		for _, ptr := range stack {
			traverse(ptr)
		}
	}

	return trace
}
