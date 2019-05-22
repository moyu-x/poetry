package main

import (
	"fmt"

	"poetry/configs"
	"poetry/tools"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	dbURL := configs.GetDatabaseUrl()

	fmt.Println(dbURL)
	db, err := gorm.Open("mysql", dbURL)
	tools.CheckErr(err)
	defer db.Close()
}
