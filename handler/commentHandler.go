package handler

import (
	"blog/model"
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发表评论 实现评论的创建功能，已认证的用户可以对文章发表评论。
func PubComment(c *gin.Context) {
	user, err := GetCurrUser(c)
	if err != nil {
		util.JSON(c, http.StatusBadRequest, err.Error(), false)
	}
	var comment model.Comment
	if errs := c.ShouldBindJSON(&comment); errs != nil {
		util.JSON(c, http.StatusBadRequest, errs.Error(), false)
	}
	comment.UserID = user.ID
	if errs := service.CreateComment(&comment); errs != nil {
		util.JSON(c, http.StatusBadRequest, errs.Error(), false)
	}
	util.JSON(c, http.StatusOK, util.PublishCommentSuccess, true)
}

// 获取某篇文章的所有评论列表
func FindCommentsByPostId(c *gin.Context) {
	var comment model.Comment
	if errs := c.ShouldBindJSON(&comment); errs != nil {
		util.JSON(c, http.StatusBadRequest, errs.Error(), nil)
	}
	comments := service.FindCommentsByPostId(comment.PostID)
	util.JSON(c, http.StatusOK, util.GetCommentSuccess, comments)
}
