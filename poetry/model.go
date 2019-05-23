package poetry

import (
	"github.com/jinzhu/gorm"
)

// Author 用来描述作者的信息
type Author struct {
	gorm.Model
	Name             string `gorm:"type:varchar(255);index" json:"name"`
	Desc             string `gorm:"type:text" json:"desc"`
	Type             string `gorm:"type:varchar(255)"`
	ShortDescription string `gorm:"type:text" json:"short_description"`
}

// JSONPoetry 诗的 JSON 结构
type JSONPoetry struct {
	Strains    []string `json:"strains"`
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Title      string   `json:"title"`
}

// Poetry 是诗歌在数据库中的映射
type Poetry struct {
	gorm.Model
	Strains    string `gorm:"type:text"`
	Author     string `gorm:"type:varchar(255);index"`
	Paragraphs string `gorm:"type:text"`
	Title      string `gorm:"type:varchar(255);index"`
	Type       string `gorm:"type:varchar(255)"`
}
