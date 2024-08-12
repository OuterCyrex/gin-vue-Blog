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

func GetArticleByCategory(pageSize int, pageNum int, cid int) ([]Article, int, int64) {
	var cate Category
	var total int64
	db.Where("id = ?", cid).First(&cate)
	if cate.ID < 0 {
		return []Article{}, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	var cateArtList []Article
	if cate.ID == 0 {
		db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cateArtList)
		db.Model(&cate).Count(&total)
		return cateArtList, errmsg.SUCCESS, total
	}
	result := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Cid = ?", cid).Find(&cateArtList)
	db.Model(&cate).Count(&total)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []Article{}, errmsg.ERROR_ARTICLE_NOT_EXIST, total
		}
		return []Article{}, errmsg.ERROR, total
	}
	return cateArtList, errmsg.SUCCESS, total
}

//查询单个文章

func SearchArticleByCid(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id=?", id).First(&article).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return Article{}, errmsg.ERROR
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Article{}, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

//查询文章列表

func GetArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var art []Article
	var total int64
	var err error
	if title != "" {
		err = db.Order("updated_at DESC").Preload("Category").Where("title LIKE ?", title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
		db.Model(&art).Where("title LIKE ?", title+"%").Count(&total)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("文章查询出错,%v", err)
			return nil, errmsg.ERROR, 0
		}
		return art, errmsg.SUCCESS, total
	} else {
		err = db.Order("updated_at DESC").Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
		db.Model(&art).Count(&total)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("文章查询出错,%v", err)
			return nil, errmsg.ERROR, 0
		}
		return art, errmsg.SUCCESS, total
	}
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
