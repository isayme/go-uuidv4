package uuidv4

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	uuid, err := Generate()

	fmt.Println(uuid, err)

	if err != nil {
		t.Error("err should be nil")
	}

	if len(uuid) != 36 {
		t.Errorf("uuid length not valid: %d", len(uuid))
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate()
	}
}
