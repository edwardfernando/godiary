package config_test

import (
	"strings"
	"testing"

	"github.com/edwardfernando/godiary/config"
	"github.com/stretchr/testify/assert"
)

func TestViperGetString(t *testing.T) {
	vp := config.NewViper()
	in := strings.NewReader(`{"foo":"bar"}`)
	vp.SetConfigType("json")
	vp.ReadConfig(in)
	assert.Equal(t, "bar", vp.GetString("foo"))
}
