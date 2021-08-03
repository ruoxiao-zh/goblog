package bootstrap

import (
	"goblog/pkg/cache"
	"goblog/pkg/config"
)

func SetupCache() {
	if config.Env("CACHE_DRIVER") == "redis" {
		cache.InitClient()
	}

}
