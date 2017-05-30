package etcd

import (
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type Etcd struct {
	Client  client.Client
	KeysAPI client.KeysAPI
	Config  client.Config
}

func New(cfg client.Config) (e *Etcd, err error) {
	e = &Etcd{}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	e.KeysAPI = client.NewKeysAPI(c)
	return e, nil
}

func (e Etcd) GetInt(key string) (int, error) {
	value, err := e.Get(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(value)
}

func (e Etcd) GetFloat(key string) (float64, error) {
	value, err := e.Get(key)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(value, 64)
}

func (e Etcd) Get(key string) (string, error) {
	resp, err := e.KeysAPI.Get(context.Background(), key, nil)
	if err != nil {
		return "", err
	}

	return resp.Node.Value, nil
}

func (e Etcd) SetOrUpdate(key, value string) error {
	_, err := e.KeysAPI.Set(context.Background(), key, value, &client.SetOptions{PrevExist: client.PrevExist})
	if err != nil {
		return err
	}
	return nil
}

func Int(value int, err error) int {
	if err != nil {
		log.WithError(err).Fatal()
	}

	return value
}

func Float(value float64, err error) float64 {
	if err != nil {
		log.WithError(err).Fatal()
	}

	return value
}

func String(value string, err error) string {
	if err != nil {
		log.WithError(err).Fatal()
	}

	return value
}
