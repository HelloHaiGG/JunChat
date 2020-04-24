package utils

import (
	"fmt"
	"testing"
)

func TestSnowFlake_GetId(t *testing.T) {
	InitSF(20)
	fmt.Println(SFIdTool.GetID())
}

func BenchmarkSnowFlake_GetId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InitSF(12)
		fmt.Println(SFIdTool.GetID())
	}
}