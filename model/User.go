package model

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gin-vue-blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(20);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(20);not null"`
	Role     int    `json:"role" type:"int"`
}

//查询用户是否存在

func CheckUser(username string) int {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//新增用户

func CreateUser(data *User) int {
	data.Password = ScryptPwd(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表

func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("用户查询出错,%v", err)
		return nil
	}
	return users
}

//编辑用户信息

func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	result := db.Model(&User{}).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	fmt.Println(result.RowsAffected)
	return errmsg.SUCCESS
}

//删除用户

func DeleteUser(id int) int {
	var users User
	result := db.Delete(&users, id)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

//密码加密
//密码加密采用的是go语言自带的scrypt库

func ScryptPwd(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 23, 4, 61, 2, 9, 92}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
