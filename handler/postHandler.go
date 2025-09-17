package handler

import (
	"blog/model"
	"blog/service"
	"blog/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建文章
func CreatePost(c *gin.Context) {
	var post model.Post
	user, err := GetCurrUser(c)
	if err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
	}
	post.UserID = user.ID
	if c.ShouldBindJSON(&post) != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
	}

	if err := service.CreatePost(&post); err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error())
	}
	util.JSON(c, http.StatusOK, util.CreatePostSuccess, true)
}

// 获取文章列表
func FindPostsByCondition(c *gin.Context) {
	var post model.Post
	user, err := GetCurrUser(c)
	if err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	post.UserID = user.ID
	if err := c.ShouldBindJSON(&post); err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	posts := service.FindPostsByCondition(&post)
	util.JSON(c, http.StatusOK, util.GetPostSuccess, posts)
}

// 修改文章
func UpdatePost(c *gin.Context) {
	var post model.Post

	if ok, err := CheckCurrUserPost(&post, c); err != nil && !ok {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
	}
	if err := service.UpdatePost(&post); err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
	}
	util.JSON(c, http.StatusOK, util.ModifyPostSuccess, true)
}

// 删除文章
func DeletePost(c *gin.Context) {
	var post model.Post
	ok, err := CheckCurrUserPost(&post, c)
	if err != nil && !ok {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
	}
	errs := service.DeletePost(post.ID)
	if errs != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
	}
	util.JSON(c, http.StatusOK, util.DelPostSuccess, true)
}

// 校验当前用户下文章
func CheckCurrUserPost(post *model.Post, c *gin.Context) (bool, error) {
	user, err := GetCurrUser(c)
	if err != nil {
		return false, err
	}
	post.UserID = user.ID
	if err := c.ShouldBindJSON(&post); err != nil {
		return false, err
	}
	currPost := service.FindPostsById(post.ID)
	if currPost == nil {
		return false, errors.New(util.PostNotExist)
	}
	if currPost.UserID != post.UserID {
		return false, errors.New(util.CurrUserNotPost)
	}
	return true, nil
}
