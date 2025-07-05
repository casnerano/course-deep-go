package data_types

import (
	"math"
	"unsafe"
)

func IsLittleEndian() bool {
	var x int16 = 0x0001
	return *(*int8)(unsafe.Pointer(&x)) == 1
}

func ToLittleEndianV1[T uint16 | uint32 | uint64](number T) T {
	octetCount := int(unsafe.Sizeof(number))

	p := unsafe.Pointer(&number)
	le := T(0)

	for i := 0; i < octetCount; i++ {
		le += T(math.Pow(math.MaxUint8+1, float64(octetCount-i-1))) * T(*(*uint8)(unsafe.Add(p, i)))
	}

	return le
}

func ToLittleEndianV2[T uint16 | uint32 | uint64](number T) T {
	octets := (*[8]byte)(unsafe.Pointer(&number))
	octetsCount := int(unsafe.Sizeof(number))

	for i := 0; i < octetsCount/2; i++ {
		octets[i], octets[octetsCount-i-1] = octets[octetsCount-i-1], octets[i]
	}

	return number
}

func ToLittleEndianV3[T uint16 | uint32 | uint64](number T) T {
	octetsCount := int(unsafe.Sizeof(number))
	octets := unsafe.Slice((*byte)(unsafe.Pointer(&number)), octetsCount)

	for i := 0; i < octetsCount/2; i++ {
		octets[i], octets[octetsCount-i-1] = octets[octetsCount-i-1], octets[i]
	}

	return number
}

func ToLittleEndianV4[T uint16 | uint32 | uint64](number T) T {
	octetsCount := int(unsafe.Sizeof(number))
	octets := unsafe.Slice((*byte)(unsafe.Pointer(&number)), octetsCount)

	for i := 0; i < octetsCount/2; i++ {
		octets[i], octets[octetsCount-i-1] = octets[octetsCount-i-1], octets[i]
	}

	return number
}
