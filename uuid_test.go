package uuidv4

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	uuid, err := Generate()

	if err != nil {
		t.Error("err should be nil")
	}

	if len(uuid) != 35 {
		t.Errorf("uuid length not valid: %d", len(uuid))
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate()
	}
}
