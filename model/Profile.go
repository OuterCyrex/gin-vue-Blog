package model

import "gin-vue-blog/utils/errmsg"

type Profile struct {
	ID     int    `gorm:"primary_key;" json:"id"`
	Name   string `gorm:"type:varchar(20);" json:"name"`
	Desc   string `gorm:"type:varchar(300);" json:"desc"`
	QQchat string `gorm:"type:varchar(20);" json:"qqchat"`
	Bili   string `gorm:"type:varchar(20);" json:"bili"`
	Github string `gorm:"type:varchar(30);" json:"github"`
	Avatar string `gorm:"type:varchar(200);" json:"avatar"`
}

func GetProfile(id int) (Profile, int) {
	var profile Profile
	err := db.Model(&profile).Where("id = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
