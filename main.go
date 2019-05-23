package main

import (
	"poetry/configs"
	"poetry/poetry"
	"poetry/tools"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// GetDatabaseConn 用于获取数据库的连接
func GetDatabaseConn() *gorm.DB {
	dbURL := configs.GetDatabaseURL()
	db, err := gorm.Open("mysql", dbURL)
	tools.CheckErr(err)

	// 设置数据库属性
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)

	return db
}

// SaveAuthors 用来保存唐朝诗人的信息
func SaveAuthors(dynasty string, authors []poetry.Author) {
	db := GetDatabaseConn()
	defer db.Close()

	if !db.HasTable(&poetry.Author{}) {
		db.CreateTable(&poetry.Author{})
	}

	for _, author := range authors {
		author.Type = dynasty
		db.NewRecord(author)
		db.Create(&author)
	}
}

// SavePoetrys 保存诗歌到数据库中
func SavePoetrys(poetrys []poetry.Poetry) {
	db := GetDatabaseConn()
	defer db.Close()

	if !db.HasTable(&poetry.Poetry{}) {
		db.CreateTable(&poetry.Poetry{})
	}

	for _, poetry := range poetrys {
		db.NewRecord(poetry)
		db.Create(&poetry)
	}
}

func main() {
	// 读取作者信息
	dynastyArray := []string{"SONG", "TANG"}
	for _, dynasty := range dynastyArray {
		authors := poetry.ReadPoetryAuthors(strings.ToLower(dynasty))
		SaveAuthors(dynasty, authors)
	}

	// 保存诗歌
	poetrys := poetry.ReadPoetry()
	SavePoetrys(poetrys)
}
