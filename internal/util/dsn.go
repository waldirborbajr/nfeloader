package util

import "github.com/waldirborbajr/nfeloader/internal/config"

func Dsn() string {
	return config.MysqlUrl(config.Cfg)
}
