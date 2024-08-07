package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `json:"title" gorm:"type:varchar(100);not null"`
	Cid      int      `json:"cid" gorm:"type:int;not null;"`
	Desc     string   `json:"desc" gorm:"type:varchar(200)"`
	Content  string   `json:"content" gorm:"type:longtext;not null"`
	Img      string   `json:"img" gorm:"type:varchar(200)"`
}
