package utils

import (
	"testing"
	"github.com/irisnet/iris-community/config"
)

func TestRegisterEmail(t *testing.T) {
	config.LoadConfiguration("../config.yml")
	RegisterEmail("XXX@qq.com","XXX","65")
}