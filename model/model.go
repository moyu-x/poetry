package model

import (
	"github.com/jinzhu/gorm"
)

// Author 用来描述作者信息
type Author struct {
	gorm.Model
	Name             string `gorm:"type:varchar(255);index"`
	Description      string `gorm:"type:text" json:"desc,description"`
	ShortDescription string `gorm:"type:text"`
	Type             string `gorm:"type:varchar(255)"`
}

// Poetry 用来描述诗歌信息，也包括宋词
type Poetry struct {
	gorm.Model
	Author  string    `gorm:"type:varchar(255)"`
	Title   string    `gorm:"type:varchar(255)"`
	Type    string    `gorm:"type:varchar(255)"`
	Content []Content `gorm:"foreignkey:PoetryContent"`
}

// Content 是诗或词的正文，或则是韵脚
type Content struct {
	gorm.Model
	Item string `gorm:"type:text"`
	Type string `gorm:"type:varchar(255)"`
}
