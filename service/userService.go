package service

import (
	"blog/model"
)

// 注册用户
func RegisterUser(user *model.User) (bool, error) {
	db := loadConfGetDb()
	db.Create(user)
	return true, nil
}

// 获取用户
func GetUserById(id uint) *model.User {
	db := loadConfGetDb()
	var user *model.User
	db.First(user, id)
	return user
}

func FindUserByCondition(kv map[string]interface{}) *[]model.User {
	db := loadConfGetDb()
	var user *[]model.User
	db.Where(kv).Find(&user)
	return user
}
