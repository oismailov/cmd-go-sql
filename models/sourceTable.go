package models

import "cmd-go-sql/config"

//finds table that ends with {name} and returns its name
func SourceTableFinder(name string) string {
	var sourceTable string
	sql := `
  SELECT TABLE_NAME AS table_name
  FROM information_schema.TABLES
  WHERE TABLE_SCHEMA = ?
  AND TABLE_NAME LIKE ?
  ORDER BY CREATE_TIME DESC
  LIMIT 1`

	db := GetDatabaseSession()
	db.Raw(sql, config.Cfg.DatabaseSettings.DatabaseName, "%"+name).Row().Scan(&sourceTable)

	return sourceTable
}
