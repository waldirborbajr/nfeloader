package util

import "github.com/waldirborbajr/nfeloader/internal/pkg/config"

func Dsn() string {
	return config.MysqlUrl(config.Cfg)
}
