package main

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func BenchmarkJsonIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(Game{
			ID:   1,
			Name: "test",
		})
	}
}

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(Game{
			ID:   1,
			Name: "test",
		})
	}
}
