package cache

import (
	"testing"

	"encoding/gob"

	log "github.com/Sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	gob.Register(&user{})
	var c Cacher = &Redis{Host: "localhost", Port: 6379, Password: "sss"}
	defer c.Close()
	c.Init()
	assert.False(t, c.Exists("sdf"))
	u := user{"sdvdxl", 27}

	c.Set("a", u)
	c.Set("name", "sdvdxl")
	var name string
	c.Get("name", &name)
	gu := user{}
	assert.True(t, c.Get("a", &gu))
	assert.True(t, "sdvdxl" == gu.Name, "test name")
	assert.Equal(t, c.GetString("name"), "sdvdxl")
	_ = c.GetString("")
}

type user struct {
	Name string
	Age  int
}
