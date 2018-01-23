package utils

import (
	"testing"
	"github.com/community/config"
)

func TestRegisterEmail(t *testing.T) {
	config.LoadConfiguration("../config.yml")
	RegisterEmail("XXX@qq.com","XXX","65")
}