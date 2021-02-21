package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// performance 性能测试

func BenchmarkConcatStringByAdd(b *testing.B) {
	elms := []string{"1", "2", "3", "4", "5"}
	// 性能测试开始
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elms {
			ret += elem
		}
	}
	// 性能测试结束
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elms := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elms {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elms := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elms {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elms := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elms {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}
