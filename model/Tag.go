package model

import (
	"errors"
	"fmt"
	"gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Tag struct {
	ID       uint     `gorm:"primary_key;auto_increment" json:"id"`
	Name     string   `gorm:"type:varchar(20);not null" json:"name"`
	Cid      int      `json:"cid" gorm:"type:int;not null;"`
	Category Category `gorm:"foreignkey:Cid"`
}

//查询单个标签

func SearchTag(id uint) (Tag, int) {
	var tag Tag
	db.Preload("Category").Where("id=?", id).First(&tag)
	if tag.ID == 0 {
		return Tag{}, errmsg.ERROR_TAG_NOT_EXIST
	}
	return tag, errmsg.SUCCESS
}

//查询标签是否存在

func CheckTag(name string) int {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return errmsg.ERROR_TAG_USED
	}
	return errmsg.SUCCESS
}

//编辑防重名

func CheckUpTag(id uint, name string) int {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID == id || tag.ID <= 0 {
		return errmsg.ERROR_TAG_USED
	}
	return errmsg.SUCCESS
}

//新增标签

func CreateTag(data *Tag) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询标签列表

func GetTag() ([]Tag, int, int64) {
	var tag []Tag
	var total int64
	var err error
	err = db.Preload("Category").Find(&tag).Error
	db.Model(&tag).Count(&total)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("标签查询出错,%v", err)
		return nil, errmsg.ERROR, 0
	}
	return tag, errmsg.SUCCESS, total
}

//通过分类查询标签

func GetTagByCategory(id uint) ([]Tag, int, int64) {
	var tag []Tag
	var total int64
	var err error
	fmt.Println(id)
	err = db.Preload("Category").Where("Cid = ?", id).Find(&tag).Error
	db.Model(&tag).Where("Cid = ?", id).Count(&total)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("标签查询出错,%v", err)
		return nil, errmsg.ERROR, 0
	}
	return tag, errmsg.SUCCESS, total
}

//编辑标签信息

func EditTag(id int, data *Tag) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	result := db.Model(&Tag{}).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_TAG_USED
	}
	fmt.Println(result.RowsAffected)
	return errmsg.SUCCESS
}

//删除标签

func DeleteTag(id int) int {
	var tag Tag
	result := db.Delete(&tag, id)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_TAG_NOT_EXIST
	}
	return errmsg.SUCCESS
}
