package lib

import (
	"github.com/ertgl/croncache/models"
)

type CacheEngine interface {
	Dependency
	Configurable
	Upsert(models.Cache) error
	Close() error
}
