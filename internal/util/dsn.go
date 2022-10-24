package util

import "github.com/waldirborbajr/nfeloader/pkg/config"

func Dsn() string {
	return config.MysqlUrl(config.Cfg)
}
