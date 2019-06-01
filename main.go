package main

import (
	"poetry/configs"
	"poetry/model"
	"poetry/poetry"
	"poetry/tools"

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

// SaveAuthors 用来保存作者相关信息
func SaveAuthors(authors []model.Author) {
	db := GetDatabaseConn()
	defer db.Close()

	if !db.HasTable(&model.Author{}) {
		db.CreateTable(&model.Author{})
	}

	for _, author := range authors {
		db.NewRecord(author)
		db.Create(&author)
	}
}

// // SavePoetrys 保存诗歌到数据库中
// func SavePoetrys(poetrys []poetry.Poetry) {
// 	db := GetDatabaseConn()
// 	defer db.Close()

// 	if !db.HasTable(&poetry.Poetry{}) {
// 		db.CreateTable(&poetry.Poetry{})
// 	}

// 	for _, poetry := range poetrys {
// 		db.NewRecord(poetry)
// 		db.Create(&poetry)
// 	}
// }

// // SaveCis 保存宋词到数据库中
// func SaveCis(cis []ci.Ci) {
// 	db := GetDatabaseConn()
// 	defer db.Close()

// 	if !db.HasTable(&ci.Ci{}) {
// 		db.CreateTable(&ci.Ci{})
// 	}

// 	for _, ci := range cis {
// 		db.NewRecord(ci)
// 		db.Create(&ci)
// 	}
// }

func main() {
	// 读取作者信息
	// dynastyArray := []string{"SONG", "TANG"}
	// floderPath := []string{"json", "ci"}
	// for _, floder := range floderPath {
	// 	if "ci" == floder {
	// 		temp := poetry.ReadPoetryAuthors("SONG", "ci")
	// 		SaveAuthors(temp)

	// 	} else {
	// 		for _, dynasty := range dynastyArray {
	// 			temp := poetry.ReadPoetryAuthors(dynasty, "json")
	// 			SaveAuthors(temp)
	// 		}
	// 	}
	// }

	poetry.ReadPoetry("TANG", "json")

	// // 保存诗歌
	// poetrys := poetry.ReadPoetry()
	// SavePoetrys(poetrys)

	// // 保存宋词
	// cis := ci.ReadCi()
	// SaveCis(cis)

	// // 保存宋朝词人
	// ciAuthors := ci.ReadCiAuthors()
	// SaveAuthors(ciAuthors)
}
