## A Light weight Go-SQL Application
This is a small software application developed in Golang. It can be run from a command prompt. It takes input parameters, interact with mysql database, and perform SQL operations.

### This application will perform following operations
1. It takes first parameter (if provided) and search a table with name ending at given value and has been most recently created. Lets call this table as 'Source' from here on. If parameter wasn't provided - it takes default value: `_ReportName`
2. The table found in step 1 is read and some transformations (listed below) are applied on it, such that the original table is left unchanged, and the transformed data is saved as a new table with name as "timestamp_ReportName_New". The newly created table will be referred to as 'Destination' in this description from here on.

### The transformations are
1. A new field will be inserted as a part of the transformation, in the Destination. This newly created field will have some random values.
2. Each record will be duplicated in the Destination, such that the duplicate record will be exactly below to the original record.

### How to use
Make sure you have GO installed on your machine, otherwise go to [How to install GO](https://golang.org/doc/install). Once Golang has been installed please go to `config/conf.json` file and change your databse settings: user, password, host, port, database name. Then cd (change directory) to your root project folder and create a build: `go build` - it will create binary file with project name `cmd-go-sql`.<br>
Now you can run application<br>
`./cmd-go-sql` - without parameter<br>
or<br>
`./cmd-go-sql -source="table name"` with parameter.<br>
You will see all information about newly created table and intermediate processes in prompt

###Update 2017-03-05
A second parameter for the command has been introduced, for the destination table name. If we run this utility again and again with the same destination table name (as a second argument to the command), it should overwrite the existing table (if exists with the same name) without raising any issue.

#### Copyright
Copyright © 2017 Infotecture Development Pty Ltd.
