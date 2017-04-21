package v1

import (
	"encoding/json"
	"fmt"
)

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

import (
	"github.com/ertgl/croncache/lib"
	"github.com/ertgl/croncache/models"
)

type CacheEngine struct {
	lib.CacheEngine
	moduleName string
	ioc        lib.IoC
	config     *Config
	connPool   *pool.Pool
}

func NewCacheEngine() lib.CacheEngine {
	ce := &CacheEngine{
		moduleName: MODULE_NAME,
		config:     NewConfig(),
	}
	return ce
}

func (ce *CacheEngine) ModuleName() string {
	return ce.moduleName
}

func (ce *CacheEngine) Initialize() error {
	var err error = nil
	onConnect := func(network, addr string) (*redis.Client, error) {
		client, err := redis.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		if len(ce.config.Password) > 0 {
			if err = client.Cmd("AUTH", ce.config.Password).Err; err != nil {
				client.Close()
				return nil, err
			}
		}
		if len(ce.config.Database) > 0 {
			if err = client.Cmd("SELECT", ce.config.Database).Err; err != nil {
				client.Close()
				return nil, err
			}
		}
		return client, nil
	}
	ce.connPool, err = pool.NewCustom(
		"tcp",
		fmt.Sprintf("%s:%d", ce.config.Host, ce.config.Port),
		ce.config.ConnectionPoolSize,
		onConnect,
	)
	return err
}

func (ce *CacheEngine) IoC() lib.IoC {
	return ce.ioc
}

func (ce *CacheEngine) SetIoC(ioc lib.IoC) {
	ce.ioc = ioc
}

func (ce *CacheEngine) ImportConfig(raw []byte) error {
	err := json.Unmarshal(raw, &ce.config)
	return err
}

func (ce *CacheEngine) ExportConfig() ([]byte, error) {
	raw, err := json.MarshalIndent(&ce.config, "", "\t")
	return raw, err
}

func (ce *CacheEngine) Upsert(cache models.Cache) error {
	conn, err := ce.connPool.Get()
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, macro := range cache.Macros {
		conn.PipeAppend(macro.Command, macro.Args...)
	}
	err = conn.PipeResp().Err
	if err != nil {
		return err
	}
	ce.connPool.Put(conn)
	return nil
}

func(ce *CacheEngine) Close() error {
	ce.connPool.Empty()
	return nil
}
