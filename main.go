package main

import (
	"cmd-go-sql/config"
	"cmd-go-sql/models"
	"fmt"
	"os"
)

var defaultTableName = "ReportName"
var tableNameToFind string

func main() {

	//load data from conf/config.json file
	config.LoadConfig()

	//connect to mysql database
	models.InitDatabaseSession()

	//read arguments from cmd
	if len(os.Args) == 1 {
		tableNameToFind = defaultTableName
		fmt.Println("you didn't provide table name. Default value is '" + defaultTableName + "'")
	} else {
		tableNameToFind = os.Args[1]
	}

	//finds table that ends with name {tableNameToFind}
	sourceTable := tableFinder(tableNameToFind)

	//If such table as {tableNameToFind} was found - it will create destination table
	if sourceTable != "" {
		fmt.Printf("founded table name is: %s\n", sourceTable)
		destinationTable := createDestinationTable(sourceTable, tableNameToFind)
		fmt.Println("successfully created destination table with name '" + destinationTable + "'")

		//If there is no such table as {tableNameToFind}, it will use {defaultTableName} as a default option
	} else if sourceTable == "" && tableNameToFind != defaultTableName {
		fmt.Println("there is no table that ends with name '" + tableNameToFind + "'. looking for a table that ends with name '" + defaultTableName + "'")
		sourceTable := tableFinder(defaultTableName)

		//looking for a table that ends with {defaultTableName}
		if sourceTable != "" {
			fmt.Printf("founded table name is: %s\n", sourceTable)
			destinationTable := createDestinationTable(sourceTable, defaultTableName)
			fmt.Println("successfully created destination table with name '" + destinationTable + "'")
		} else {
			fmt.Println("table was not found")
		}
	} else {
		fmt.Println("table was not found")
	}

}

func tableFinder(tableName string) string {
	if tableName != defaultTableName {
		fmt.Println("looking for a table that ends with '" + tableName + "'")
	}

	return models.SourceTableFinder(tableName)
}

func createDestinationTable(sourceTable string, tableNameToFind string) string {
	return models.CreateDestinationTable(sourceTable, tableNameToFind)
}
