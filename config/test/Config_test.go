package test

import (
	c "dimples-api/config"
	"testing"
)

func TestParseConfig(t *testing.T) {
	cfg := c.ParseConfig("../../Haste.config")
	t.Log("UnMarshall obj", cfg)
}