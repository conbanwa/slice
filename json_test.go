package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBodyCut(t *testing.T) {
	assert.Equal(t, BodyCut(longBytes),
		"package lap\n\nimport (\n\t\"sync/atomic\"\n\t\"time\"\n\n\t\"github.com/conbanwa/logs\"\n)\n\ntype Elapse []int64\nty")
	assert.Equal(t, BodyCut([]byte("package lap\n\nimport (\n\t\"sync/atomic"), 7), "package")

}

// 测试标准转换string()性能
func BenchmarkNormalBytes2String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(longBytes)
	}
}

// 测试强转换[]byte到string性能
func BenchmarkBytes2String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Bytes2String(longBytes)
	}
}

func FuzzStringBytes(f *testing.F) {
	f.Add(longBytes)
	f.Fuzz(func(t *testing.T, a []byte) {
		assert.Equal(t, a, String2Bytes(Bytes2String(a)))
	})
}
