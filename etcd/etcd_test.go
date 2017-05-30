package etcd

import (
	"testing"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/stretchr/testify/assert"
)

func TestEtcd(t *testing.T) {
	etcdCfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}

	e, err := New(etcdCfg)
	assert.Nil(t, err)

	assert.Nil(t, e.SetOrUpdate("/test/a/b", "value"))
	assert.Equal(t, String(e.Get("/test/a/b")), "value")
}
