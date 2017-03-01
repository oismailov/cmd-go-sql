## A Light weight Go-SQL Application
This is a small software application developed in Golang. It can be run from a command prompt. It takes input parameters, interact with mysql database, and perform SQL operations.

### This application will perform following operations
1. It finds such a table with name ending at `_ReportName` and has been most recently created. Lets call this table as 'Source' from here on.
2. The table found in step 1 is read and some transformations (listed below) are applied on it, such that the original table is left unchanged, and the transformed data is saved as a new table with name as "timestamp_ReportName_New". The newly created table will be referred to as 'Destination' in this description from here on.

### The transformations are
1. A new field will be inserted as a part of the transformation, in the Destination. This newly created field will have some random values.
2. Each record will be duplicated in the Destination, such that the duplicate record will be exactly below to the original record.

### How to use
Make sure you have GO installed on your machine, otherwise go to [How to install GO](https://golang.org/doc/install)
CD (change directory) to your root project folder and ceate a build: `go build` - it will create binary file with project name `cmd-go-sql`
Now you can run application
`./cmd-go-sql` - without parameter
or
`./cmd-go-sql table_name` with parameter
you will see all information about newly created table and intermediate processes in prompt

#### Copyright
Copyright © 2017 Infotecture Development Pty Ltd.
