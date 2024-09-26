package dict

import (
	"testing"
)

func BenchmarkGetDict(b *testing.B) {
	// Run the GetDict function b.N times to benchmark it
	for i := 0; i < b.N; i++ {
		GetDict()
	}
}

func BenchmarkGetDictBuffered(b *testing.B) {
	// Run the GetDict function b.N times with buffered output to benchmark it
	for i := 0; i < b.N; i++ {
		GetDictBuffer()
	}
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomWord()
	}
}

//
// func BenchmarkRandom1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		GetRandomWord1()
// 	}
// }
