package create

import (
	"testing"
)

func TestConfig_GetConf(t *testing.T) {
	GetConf().Key = "key"
	GetConf().Val = "val"
}
