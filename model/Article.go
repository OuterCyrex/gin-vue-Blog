package model

import (
	"errors"
	"fmt"
	"gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `json:"title" gorm:"type:varchar(100);not null"`
	Cid      int      `json:"cid" gorm:"type:int;not null;"`
	Desc     string   `json:"desc" gorm:"type:varchar(200)"`
	Content  string   `json:"content" gorm:"type:longtext;not null"`
	Img      string   `json:"img" gorm:"type:varchar(200)"`
}

//查询分类下的所有文章

//查询单个文章

//查询文章列表

func GetArticle(pageSize int, pageNum int) []Article {
	var art []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("文章查询出错,%v", err)
		return nil
	}
	return art
}

//添加文章

func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑文章

func EditArticle(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	result := db.Model(&Article{}).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	fmt.Println(result.RowsAffected)
	return errmsg.SUCCESS
}

//删除文章

func DeleteArticle(id int) int {
	var art Article
	result := db.Delete(&art, id)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return errmsg.SUCCESS
}
