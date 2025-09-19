package handler

import (
	"blog/model"
	"blog/service"
	"blog/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 创建文章
func CreatePost(c *gin.Context) {
	var post model.Post
	user, err := GetCurrUser(c)
	if err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
	}
	post.UserID = user.ID
	if errs := c.ShouldBindBodyWithJSON(&post); errs != nil {
		util.JSON(c, http.StatusBadRequest, util.FilterEnglish(errs), false)
		return
	}

	if errs := service.CreatePost(&post); errs != nil {
		util.JSON(c, http.StatusBadRequest, errs.Error())
		return
	}
	util.JSON(c, http.StatusOK, util.CreatePostSuccess, true)
}

// 获取文章列表
func FindPostsByCondition(c *gin.Context) {
	var post model.Post
	user, err := GetCurrUser(c)
	if err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	post.UserID = user.ID
	posts := service.FindPostsByCondition(&post)
	util.JSON(c, http.StatusOK, util.GetPostSuccess, posts)
}

// 修改文章
func UpdatePost(c *gin.Context) {
	var post model.Post

	if ok, err := CheckCurrUserPost(&post, c); err != nil && !ok {
		util.JSON(c, http.StatusBadRequest, util.FilterEnglish(err), false)
		return
	}
	if err := service.UpdatePost(&post); err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
		return
	}
	util.JSON(c, http.StatusOK, util.ModifyPostSuccess, true)
}

// 删除文章
func DeletePost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	errs := service.DeletePost(id)
	if errs != nil {
		util.JSON(c, http.StatusBadRequest, errs.Error(), false)
		return
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
	if err := c.ShouldBindJSON(post); err != nil {
		return false, err
	}
	currPost := service.FindPostsById(post.ID)
	if currPost.ID == 0 {
		return false, errors.New(util.PostNotExist)
	}
	if currPost.UserID != post.UserID {
		return false, errors.New(util.CurrUserNotPost)
	}
	if post.Title == "" {
		post.Title = currPost.Title
	}
	if post.Content == "" {
		post.Content = currPost.Content
	}
	post.CreatedAt = currPost.CreatedAt
	return true, nil
}
