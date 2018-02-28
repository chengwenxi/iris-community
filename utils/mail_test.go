package utils

import (
	"testing"
	"github.com/irisnet/iris-community/config"
)

func TestRegisterEmail(t *testing.T) {
	config.LoadConfiguration("../config.yml")
	RegisterEmail("760329367@qq.com", "1", "65")
}
