package utils

import (
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	s := "Hello world"
	key := "aaaaaaaaaaaaaaaa"
	bs := []byte(s)
	s, err := AesEncrypt(bs, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}

func TestAesDecrypt(t *testing.T) {
	body := "9Nab8FdactzmcqT3LrV6HA=="
	key := "aaaaaaaaaaaaaaaa"
	s,err := AesDecrypt(body, key)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(s)
}
