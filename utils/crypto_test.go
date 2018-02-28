package utils

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Printf(Md5("123456"))
}

func TestSha1s(t *testing.T) {
	fmt.Printf(Sha1s("123456"))
}

func TestRandomInfo(t *testing.T) {
	fmt.Printf(RandomInfo(6))
}

func TestIntTo52(t *testing.T) {
	fmt.Printf(IntTo52(6, 5000145))
}
