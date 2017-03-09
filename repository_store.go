package croncache

import (
	"github.com/ertgl/croncache/lib"
	"github.com/ertgl/croncache/repositories"
)

var (
	iocRepository         lib.IoCRepository         = repositories.NewIoCRepository()
	nodeRepository        lib.NodeRepository        = repositories.NewNodeRepository()
	executerRepository    lib.ExecuterRepository    = repositories.NewExecuterRepository()
	taskManagerRepository lib.TaskManagerRepository = repositories.NewTaskManagerRepository()
	cacheEngineRepository lib.CacheEngineRepository = repositories.NewCacheEngineRepository()
)

func IoCRepository() lib.IoCRepository {
	return iocRepository
}

func NodeRepository() lib.NodeRepository {
	return nodeRepository
}

func ExecuterRepository() lib.ExecuterRepository {
	return executerRepository
}

func TaskManagerRepository() lib.TaskManagerRepository {
	return taskManagerRepository
}

func CacheEngineRepository() lib.CacheEngineRepository {
	return cacheEngineRepository
}
