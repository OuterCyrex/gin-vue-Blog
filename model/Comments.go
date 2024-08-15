package model

import (
	"errors"
	"gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	User     User    `gorm:"foreignkey:Uid"`
	Article  Article `gorm:"foreignkey:Aid"`
	Uid      int     `json:"uid" gorm:"type:int;not null;"`
	Aid      int     `json:"aid" gorm:"type:int;not null;"`
	Comments string  `json:"comments" gorm:"type:text;not null;"`
}

//查询文章的所有评论

func GetCommentsByArticle(pageSize int, pageNum int, aid int) ([]Comments, int, int64) {
	var art Article
	var total int64
	result := db.Where("id = ?", aid).First(&art)
	var commentArt []Comments
	if result.RowsAffected <= 0 {
		return []Comments{}, errmsg.ERROR_ARTICLE_NOT_EXIST, 0
	}
	result = db.Order("updated_at DESC").Preload("Article").Preload("User").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Aid = ?", aid).Find(&commentArt)
	db.Model(&commentArt).Where("Aid = ?", aid).Count(&total)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []Comments{}, errmsg.ERROR_COMMENT_NOT_EXIST, total
		}
		return []Comments{}, errmsg.ERROR, total
	}
	if result.RowsAffected == 0 {
		return []Comments{}, errmsg.ERROR_COMMENT_NOT_EXIST, total
	}
	return commentArt, errmsg.SUCCESS, total
}

func CreateComment(data *Comments) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
