package main

import (
	"fmt"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbSourceName := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
	dbIns, err := xorm.NewEngine("mysql", dbSourceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	dbIns.ShowSQL(true)
	dbIns.Logger().SetLevel(core.LOG_DEBUG)
	dbIns.SetMaxOpenConns(20)
	dbIns.SetMaxIdleConns(20)
	sql := "SELECT column_name, column_comment FROM information_schema.columns WHERE table_name = ?"
	columnMap := dbIns.SQL(sql, "student").Query()
	if columnMap.Error != nil {
		fmt.Println(columnMap.Error)
		return
	}
	columnResults := columnMap.Results
	fmt.Printf("%+v",columnResults)

	dataMap := dbIns.SQL("SELECT * FROM student").Query()
	if dataMap.Error != nil{
		fmt.Println(dataMap.Error)
		return
	}
	fmt.Println()
	fmt.Printf("%+v",dataMap.Results)
}
