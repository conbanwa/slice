package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var longBytes = String2Bytes("package lap\n\nimport (\n\t\"sync/atomic\"\n\t\"time\"\n\n\t\"github.com/conbanwa/logs\"\n)\n\ntype Elapse []int64\ntyruct {`package lap\n\nimport (\n\t\"sync/atomic\"\n\t\"time\"\n\n\t\"github.com/conbanwa/logs\"\n)\n\ntype Elapse []int64\ntyruct {`package lap\n\nimport (\n\t\"sync/atomic\"\n\t\"time\"\n\n\t\"github.com/conbanwa/logs\"\n)\n\ntype Elapse []int64\ntyruct {`package lap\n\nimport (\n\t\"sync/atomic\"\n\t\"time\"\n\n\t\"github.com/conbanwa/logs\"\n)\n\ntype Elapse []int64\ntyruct {`package lap\n\nimport (\n\t\"sync/atomic\"\n\t\"time\"\n\n\t\"github.com/conbanwa/logs\"\n)\n\ntype Elapse []int64\ntyruct {`package lap\n\nimport (\n\t\"sync/atomic\"")

func TestSlice(t *testing.T) {
	assert.Equal(t, string(Slice([]byte(`1234567890`), 2)), "34567890")
	assert.Equal(t, string(Slice([]byte(`1234567890`), -2)), "90")
	assert.Equal(t, string(Slice([]byte(`1234567890`), 0, -4)), "123456")
	assert.Equal(t, string(Slice([]byte(`1234567890`), 2, -4)), "3456")
	assert.Equal(t, string(Slice([]byte(`1234567890`), -5, 4)), "")
	assert.Equal(t, string(Slice([]byte(`1234567890`), 6, 99)), "7890")
	assert.Equal(t, string(Slice([]byte(`1234567890`), -5, 99)), "67890")
	assert.Equal(t, string(Slice([]byte(`1234567890`), -99, 99)), "1234567890")
	assert.Equal(t, BodyCut(longBytes),
		"package lap\n\nimport (\n\t\"sync/atomic\"\n\t\"time\"\n\n\t\"github.com/conbanwa/logs\"\n)\n\ntype Elapse []int64\nty")
	assert.Equal(t, BodyCut([]byte("package lap\n\nimport (\n\t\"sync/atomic"), 7), "package")
	assert.Equal(t, Cut(longBytes),
		String2Bytes("package lap\n\nimport (\n\t\"sync/atomic\"\n\t\"time\"\n\n\t\"github.com/conbanwa/logs\"\n)\n\ntype Elapse []int64\nty"))
	assert.Equal(t, Cut([]byte("package lap\n\nimport (\n\t\"sync/atomic"), 7), String2Bytes("package"))
}

func unique[T comparable](sl *[]T) {
	*sl = MapKeys(ToMap(*sl))
}

func TestUnique(t *testing.T) {
	type args[T comparable] *[]T
	type testCase[T comparable] struct {
		name string
		args args[T]
	}
	b := []byte(`0123456789卍0123456789`)
	tests := []testCase[byte]{
		{
			name: "1",
			args: args[byte](&b),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.args
			t.Log(a)
			unique(a)
			t.Log(a)
		})
	}
}

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unique(&longBytes)
	}
}
func BenchmarkUniqueOrdered(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqSlice(longBytes)
	}
}

func FuzzUnique(f *testing.F) {
	f.Add(longBytes)
	f.Fuzz(func(t *testing.T, b []byte) {
		testToMap(t, b)
	})
}

func testToMap(t *testing.T, b []byte) {
	a := b
	assert.False(t, &a == &b)
	unique(&a)
	assert.Equal(t, ToMap(UniqSlice(b)), ToMap(a))
}
