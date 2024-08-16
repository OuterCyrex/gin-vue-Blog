package model

import (
	"errors"
	"fmt"
	"gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type FriendLink struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	Link string `gorm:"size:255;not null" json:"link"`
}

func GetFriendLink(pageSize int, pageNum int) ([]FriendLink, int, int64) {
	var friend []FriendLink
	var total int64
	var err error
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&friend).Error
	db.Model(&friend).Count(&total)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("分类查询出错,%v", err)
		return nil, errmsg.ERROR, 0
	}
	return friend, errmsg.SUCCESS, total
}

func CreateFriendLink(data *FriendLink) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteFriendLink(id int) int {
	var friend FriendLink
	result := db.Delete(&friend, id)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return errmsg.SUCCESS
}
