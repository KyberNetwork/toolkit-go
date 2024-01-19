package redis

import "time"

type Config struct {
	Addrs            []string      `mapstructure:"addrs" json:"addrs"`
	MasterName       string        `mapstructure:"master_name" json:"masterName"`
	DBNumber         int           `mapstructure:"db_number" json:"dbNumber"`
	Username         string        `mapstructure:"username" json:"username"`
	Password         string        `mapstructure:"password" json:"-"`
	SentinelUsername string        `mapstructure:"sentinel_username" json:"sentinelUsername"`
	SentinelPassword string        `mapstructure:"sentinel_password" json:"-"`
	Prefix           string        `mapstructure:"prefix" json:"prefix"`
	Separator        string        `mapstructure:"separator" json:"separator"`
	ReadTimeout      time.Duration `mapstructure:"read_timeout" json:"readTimeout"`
	WriteTimeout     time.Duration `mapstructure:"write_timeout" json:"writeTimeout"`
}
