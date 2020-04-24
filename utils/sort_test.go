package utils

import (
	"fmt"
	"testing"
)

func TestSortStr(t *testing.T) {
	s := []string{"as","av","ac","bs","sdf","sdfw"}
	fmt.Println(SortStr(s,0))
}

func BenchmarkSortStr(b *testing.B) {
	s := []string{"as","av","ac","bs","sdf","sdfw"}
	for i := 0; i < b.N; i++ {
		SortStr(s,-1)
	}
}
