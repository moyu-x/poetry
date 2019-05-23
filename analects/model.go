package analects

import (
	"github.com/jinzhu/gorm"
)

// JSONAnalects 是 JSON 的文件格式
type JSONAnalects struct {
    Chapter string `json:"chapter"`
    Paragraphs []string `json:"paragraphs"`
}

// Analects 是论语的数据库结构
type Analects struct {
    gorm.Model
    Chapter string `gorm:"type:varchar(255)"`
    Paragraphs string `gorm:"type:text"`
}
