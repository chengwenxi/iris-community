package utils
import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Printf(Md5("123456"))
}