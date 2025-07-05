package data_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}
)

func TestToLittleEndianV1(t *testing.T) {
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianV1(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestToLittleEndianV2(t *testing.T) {
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianV2(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestToLittleEndianV3(t *testing.T) {
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianV3(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func BenchmarkToLittleEndianV1_Uint64(b *testing.B) {
	var num uint64 = 0x0102030405060708
	for i := 0; i < b.N; i++ {
		_ = ToLittleEndianV1(num)
	}
}

func BenchmarkToLittleEndianV2_Uint64(b *testing.B) {
	var num uint64 = 0x0102030405060708
	for i := 0; i < b.N; i++ {
		_ = ToLittleEndianV2(num)
	}
}

func BenchmarkToLittleEndianV3_Uint64(b *testing.B) {
	var num uint64 = 0x0102030405060708
	for i := 0; i < b.N; i++ {
		_ = ToLittleEndianV3(num)
	}
}
