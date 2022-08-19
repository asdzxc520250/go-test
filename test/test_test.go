package test

import "testing"

func BenchmarkPrint(b *testing.B) {
	t := &Test{}
	for i := 0; i < b.N; i++ {
		t.Print("test")
	}
}
