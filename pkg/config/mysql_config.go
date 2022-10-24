package config

import (
	"fmt"

	"github.com/waldirborbajr/nfeloader/internal/entity"
)

func MysqlUrl(dbConfig *entity.NFeConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.DatabaseUsr,
		dbConfig.DatabasePwd,
		dbConfig.DatabaseHost,
		dbConfig.DatabaseDbName,
	)
}
