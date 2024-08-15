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
	Username string `json:"username" gorm:"type:varchar(20);not null" validate:"required,min=4,max=12" label:"用户名"`
	Password string `json:"password" gorm:"type:varchar(20);not null" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `json:"role" gorm:"type:int;DEFAULT:2" validate:"required,gte=1" label:"权限"`
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

//通过名称查询用户信息

func CheckUserByName(username string) (User, int) {
	var users User
	db.Where("username = ?", username).First(&users)
	if users.ID <= 0 {
		return users, errmsg.ERROR_USER_NOT_EXIST
	}
	return users, errmsg.SUCCESS
}

//判断重复用户名是否为新增用户

func EditCheck(id int, username string) int {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID == uint(id) || user.ID <= 0 {
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_USERNAME_USED
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

// 查询单个用户

func GetUser(id int) (User, int) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

//查询用户列表

func GetUsers(username string, pageSize int, pageNum int) ([]User, int, int64) {
	var users []User
	var total int64
	if username == "" {
		db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Count(&total)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("用户查询出错,%v", err)
			return nil, errmsg.ERROR, 0
		}
		return users, errmsg.SUCCESS, total
	} else {
		db.Where("username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where("username LIKE ?", username+"%").Count(&total)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("用户查询出错,%v", err)
			return nil, errmsg.ERROR, 0
		}
		return users, errmsg.SUCCESS, total
	}
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

//更新用户密码

func EditPsw(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["password"] = ScryptPwd(data.Password)
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

//登录验证

func CheckLogin(username, password string) int {
	var users User
	db.Where("username = ?", username).First(&users)
	// gorm未查询到时返回默认空值
	if users.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPwd(password) != users.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if users.Role != 1 {
		return errmsg.ERROR_NO_PRIVILIGE
	}
	return errmsg.SUCCESS
}

func UserLogin(username, password string) int {
	var users User
	db.Where("username = ?", username).First(&users)
	// gorm未查询到时返回默认空值
	if users.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPwd(password) != users.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}
