package ci

import (
	"github.com/jinzhu/gorm"
)

// JSONCi 是宋词的 JSON 的结构
type JSONCi struct {
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Rhythmic   string   `json:"rhythmic"`
}

// Ci 是宋词的数据库结构
type Ci struct {
	gorm.Model
	Author     string `gorm:"type:varchar(255)"`
	Paragraphs string `gorm:"type:text"`
	Rhythmic   string `gorm:"type:varchar(255)"`
}
