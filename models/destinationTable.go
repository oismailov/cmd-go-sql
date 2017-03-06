package models

import (
	"strconv"
	"time"
)

//creates copy of table {name} and give it a name {name}_New
func CreateDestinationTable(name string, tableNameToFind string) string {
	unixTime := time.Now().Unix()
	t := strconv.FormatInt(unixTime, 10)
	newTableName := t + "_" + tableNameToFind + "_New"

	transformationTable := newTableName + "_transformation"

	createTransformationTable(transformationTable, newTableName, name)

	return newTableName
}

/*
  1. create copy of source table
  2. duplicate all records: duplicate record is exactly below to the original record.
  3. add field {randomFieldName} with 'copy' value only for duplicated records
*/
func createTransformationTable(transformationTable string, newTableName string, oldTableName string) {
	db := GetDatabaseSession()
	triggerName := "transformation_flow"
	randomFieldName := "random_data"
	transformationTableFinal := transformationTable + "_final"

	//copy all records from source table
	db.Exec("CREATE TABLE `" + transformationTable + "` AS SELECT * FROM `" + oldTableName + "`")

	//duplicate all records from source table
	db.Exec("INSERT INTO `" + transformationTable + "` SELECT * FROM `" + oldTableName + "`")

	//add temporary id_temp field to recognize records (for order by purposes etc.)
	db.Exec("ALTER TABLE `" + transformationTable + "` ADD COLUMN id_temp INT(11) NOT NULL")

	//create trigger on {transformationTable} to fill id_temp with ids
	db.Exec("CREATE TRIGGER `" + triggerName + "` before UPDATE ON `" + transformationTable + "`" +
		" FOR EACH ROW" +
		" BEGIN" +
		" SET @num = (SELECT MAX(id_temp) FROM `" + transformationTable + "`);" +
		" IF NEW.id_temp  = 0 THEN" +
		" SET NEW.id_temp = @num + 1;" +
		" END IF;" +
		" END;")

	//get number of records in source table
	var sourceTableRecordsCount string
	db.Raw("SELECT COUNT(*) AS cnt FROM `" + oldTableName + "`").Row().Scan(&sourceTableRecordsCount)

	//update half a table with 0 value for id_temp field
	db.Exec("UPDATE `" + transformationTable + "` SET id_temp = 0 LIMIT " + sourceTableRecordsCount)

	//remove duplicates from {transformationTable}
	db.Exec("DELETE FROM `" + transformationTable + "` WHERE id_temp = 0")

	//create final {transformationTableFinal}
	db.Exec("CREATE TABLE `" + transformationTableFinal + "` AS SELECT * FROM `" + transformationTable + "`")

	//add duplicte records to {transformationTableFinal}
	db.Exec("INSERT INTO `" + transformationTableFinal + "` SELECT  * FROM `" + transformationTable + "`")

	//add random_data field
	db.Exec("ALTER TABLE `" + transformationTableFinal + "` ADD COLUMN " + randomFieldName + " VARCHAR(16) DEFAULT ''")

	//fill random field with data
	db.Exec("UPDATE `" + transformationTableFinal + "` SET " + randomFieldName + " = 'copy' LIMIT " + sourceTableRecordsCount)

	//create destination table
	db.Exec("CREATE TABLE `" + newTableName + "` SELECT * FROM `" + transformationTableFinal + "` ORDER BY id_temp, " + randomFieldName)

	//drop id_temp field
	db.Exec("ALTER TABLE `" + newTableName + "` DROP COLUMN id_temp")

	//clear transformation tables
	db.Exec("DROP TABLE `" + transformationTableFinal + "`")
	db.Exec("DROP TABLE `" + transformationTable + "`")

	//drop trigger {triggerName}
	db.Exec("DROP TRIGGER " + triggerName)
}
