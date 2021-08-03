package config

import "goblog/pkg/config"

func init() {
	config.Add("cache", config.StrMap{

		// redis 连接信息
		"redis": map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"password": config.Env("REDIS_PASSWORD", ""),
			"port":     config.Env("REDIS_PORT", "3306"),
			"database": config.Env("REDIS_CACHE_DB", 0),
		},
	})
}
