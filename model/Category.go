package model

import (
	"errors"
	"fmt"
	"gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询单个分类

func SearchCategory(id uint) (Category, int) {
	var cate Category
	db.Where("id=?", id).First(&cate)
	if cate.ID == 0 {
		return Category{}, errmsg.ERROR_CATEGORY_USED
	}
	return cate, errmsg.SUCCESS
}

//查询分类是否存在

func CheckCategory(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS
}

//编辑防重名

func CheckUpCategory(id uint, name string) int {
	var cate Category
	fmt.Println(id, name)
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID == id || cate.ID <= 0 {
		return errmsg.ERROR_CATEGORY_NOT_EXIST
	}
	return errmsg.SUCCESS
}

//新增分类

func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类列表

func GetCategory(name string, pageSize int, pageNum int) ([]Category, int, int64) {
	var cate []Category
	var total int64
	var err error
	if name != "" {
		err = db.Where("name Like ?", name+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
		db.Model(&cate).Where("name Like ?", name+"%").Count(&total)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("分类查询出错,%v", err)
			return nil, errmsg.ERROR, 0
		}
		return cate, errmsg.SUCCESS, total
	} else {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
		db.Model(&cate).Count(&total)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("分类查询出错,%v", err)
			return nil, errmsg.ERROR, 0
		}
		return cate, errmsg.SUCCESS, total
	}
}

//编辑分类信息

func EditCategory(id int, data *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	result := db.Model(&Category{}).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	fmt.Println(result.RowsAffected)
	return errmsg.SUCCESS
}

//删除分类

func DeleteCategory(id int) int {
	var cate Category
	result := db.Delete(&cate, id)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_CATEGORY_NOT_EXIST
	}
	return errmsg.SUCCESS
}
