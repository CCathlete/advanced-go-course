package main

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

func MyFunc() *[]byte {
	data := make([]byte, 1024)
	rand.Read(data)
	preResult := sha256.Sum256(data)
	result := preResult[:]
	return &result
}

func BenchmarkMyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunc()
	}
}
