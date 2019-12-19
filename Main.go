package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type DBStruct struct {
	ColName string
	IsNull string
	ColType string
	ColKey string
	ColComment string
}

type WordStruct struct {
	Title string
	Field string
}

// root 192.168.1.21 Kitche931743 custominfo output.doc
func main() {
	args := os.Args
	dataBaseUserName := args[1]
	dataBaseHostName := args[2]
	dataBasePassword := args[3]
	dataBaseName := args[4]
	docName := args[5]

	var wordUtil WordUtil
	columns := []WordStruct{
		WordStruct{
			Title: "字段名",
			Field: "ColName",
		},
		WordStruct{
			Title: "是否为空",
			Field: "IsNull",
		},
		WordStruct{
			Title: "字段类型",
			Field: "ColType",
		},

		WordStruct{
			Title: "数据库键",
			Field: "ColKey",
		},

		WordStruct{
			Title: "描述",
			Field: "ColComment",
		},

	}

	//databaseDocName := args[4]
	db,err := ConnectDatabase(dataBaseUserName,dataBaseHostName,dataBasePassword,dataBaseName)
	if err == nil {
		defer db.Close()
		fmt.Print("connect success")
		rows, err := db.Query("show tables")
		if err == nil {
			infoSchema,err := ConnectInforationSchema(dataBaseUserName,dataBaseHostName,dataBasePassword)
			if err == nil {
				defer infoSchema.Close()

				var dataBaseInfo map[string] []DBStruct
				dataBaseInfo = make(map[string] []DBStruct)
				for rows.Next() {
					var data []DBStruct
					var tableName string
					if err := rows.Scan(&tableName); err != nil {
						log.Fatal(err)
					} else {
						fmt.Printf("table %s\n", tableName)
						sqlstr := "select COLUMN_NAME,IS_NULLABLE,COLUMN_TYPE,COLUMN_KEY,COLUMN_COMMENT from COLUMNS where table_schema=? and table_name=? "
						cols,err := infoSchema.Query(sqlstr,dataBaseName,tableName)
						if err == nil {
							for cols.Next() {
								var colName,isNullAble,columnType,columnKey,columnComment string
								if err := cols.Scan(&colName,&isNullAble,&columnType,&columnKey,&columnComment); err != nil {
									log.Fatal(err)
								} else {
									fmt.Printf("colName:%s,columnComment:%s\n",colName,columnComment)
									var dbstruct DBStruct
									dbstruct.ColName = colName
									dbstruct.IsNull = isNullAble
									dbstruct.ColKey = columnKey
									dbstruct.ColType = columnType
									dbstruct.ColComment = columnComment
									data = append(data,dbstruct)
								}
							}
						}
					}
					dataBaseInfo[tableName] = data
				}
				wordUtil.WriteTableInfo(columns,dataBaseInfo,docName)

			}

		}

	}
}



func ConnectInforationSchema(dataBaseUserName string,
	dataBaseHostName string,
	dataBasePassword string) (*sql.DB,error) {

	var config mysql.Config
	config.DBName = "information_schema"
	config.User = dataBaseUserName
	config.Passwd = dataBasePassword
	hostName := fmt.Sprintf("tcp(%s:3306)",dataBaseHostName)
	config.Net = hostName
	print(config.FormatDSN())
	db, err := sql.Open("mysql", config.FormatDSN())
	return db,err
}


//username:password@protocol(address)/dbname?param=value
func ConnectDatabase(dataBaseUserName string,
		dataBaseHostName string,
		dataBasePassword string,
		dataBaseName string) (*sql.DB,error) {

	var config mysql.Config
	config.DBName = dataBaseName
	config.User = dataBaseUserName
	config.Passwd = dataBasePassword
	hostName := fmt.Sprintf("tcp(%s:3306)",dataBaseHostName)
	config.Net = hostName
	print(config.FormatDSN())
	db, err := sql.Open("mysql", config.FormatDSN())
	return db,err
}

