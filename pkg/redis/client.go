package redis

import "github.com/redis/go-redis/v9"

func New(cfg Config) redis.UniversalClient {
	if cfg.MasterName != "" {
		return redis.NewFailoverClusterClient(&redis.FailoverOptions{
			MasterName:       cfg.MasterName,
			SentinelAddrs:    cfg.Addrs,
			DB:               cfg.DBNumber,
			Username:         cfg.Username,
			Password:         cfg.Password,
			SentinelUsername: cfg.SentinelUsername,
			SentinelPassword: cfg.SentinelPassword,
			ReadTimeout:      cfg.ReadTimeout,
			WriteTimeout:     cfg.WriteTimeout,
			RouteRandomly:    cfg.RouteRandomly,
			ReplicaOnly:      cfg.ReplicaOnly,
		})
	}

	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            cfg.Addrs,
		MasterName:       cfg.MasterName,
		DB:               cfg.DBNumber,
		Username:         cfg.Username,
		Password:         cfg.Password,
		SentinelUsername: cfg.SentinelUsername,
		SentinelPassword: cfg.SentinelPassword,
		ReadTimeout:      cfg.ReadTimeout,
		WriteTimeout:     cfg.WriteTimeout,
	})
}
