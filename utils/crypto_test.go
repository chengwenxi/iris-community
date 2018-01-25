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

func Test(t *testing.T){
	b := 'a' + 1
	fmt.Printf(string(b))
}