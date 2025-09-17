package handler

import (
	"blog/model"
	"blog/service"
	"blog/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// 注册用户
func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.JSON(c, http.StatusBadRequest, util.TypeTransferFail)
		return
	}
	encryptPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		util.JSON(c, http.StatusBadRequest, util.PwdCryptFail)
		return
	}

	user.Password = string(encryptPwd)
	result, err := service.RegisterUser(&user)
	if err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error())
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", util.GenJWT(&user)))
	util.JSON(c, http.StatusOK, util.RegisterSuccess, result)
}

// 登录
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error())
		return
	}
	users := service.FindUserByCondition(map[string]interface{}{"username": user.UserName})
	if len(*users) == 0 {
		util.JSON(c, http.StatusBadRequest, util.UserNotExist)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte((*users)[0].Password), []byte(user.Password)); err != nil {
		util.JSON(c, http.StatusBadRequest, util.PwdError)
		return
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", util.GenJWT(&user)))
	util.JSON(c, http.StatusOK, util.LoginSuccess)
}

// 获取当前用户信息
func GetCurrUser(c *gin.Context) (*model.User, error) {
	user, ok := c.Get("user")
	if !ok {
		return nil, errors.New(util.GetCurrentUserFail)
	}
	userModel, ok := user.(model.User)
	if !ok {
		return nil, errors.New(util.TypeTransferFail)
	}
	return &userModel, nil
}
