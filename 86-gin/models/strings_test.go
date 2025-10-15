package models_test

import (
	"demo/models"
	"testing"
)

func TestStringReverse(t *testing.T) {
	str := "abcdef"
	result := models.ReverseStr(str)
	expected := "fedcba"

	if result != expected {
		t.Fail()
	}

}

func BenchmarkContact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		models.ConcatStr()
	}
}

func BenchmarkContactStrBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		models.ConcatStrBuilder()
	}
}
