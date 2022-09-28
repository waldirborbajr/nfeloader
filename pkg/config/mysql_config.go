package config

import (
	"fmt"
)

func MysqlUrl(dbConfig *NFeConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.DatabaseUsr,
		dbConfig.DatabasePwd,
		dbConfig.DatabaseHost,
		dbConfig.DatabaseDbName,
	)
}
