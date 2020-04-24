package utils

import (
	"fmt"
	"testing"
)

func Test_randomLen(t *testing.T) {
	t.Log(randomLen())
}
func Test_randomColor(t *testing.T) {
	t.Log(randomColor())
}
func TestCreateImageCode(t *testing.T)  {
	id,code := CreateImageCode()
	fmt.Println("id:",id)
	fmt.Println("code:",code)
}
func BenchmarkCreateImageCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateImageCode()
	}
}
