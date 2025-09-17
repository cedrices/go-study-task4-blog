package handler

import "github.com/gin-gonic/gin"

func InitController(api *gin.RouterGroup) {
	//注册
	api.POST("user/register", RegisterUser)
	//登录
	api.POST("user/login", Login)
	//创建文章
	api.POST("post/create", CreatePost)
	//获取文章列表
	api.GET("post/list", FindPostsByCondition)
	//修改文章
	api.PUT("post/update", UpdatePost)
	//删除文章
	api.POST("post/delete", DeletePost)
	//发布评论
	api.POST("comment/publish", PubComment)
	//获取评论
	api.GET("comment/list", FindCommentsByPostId)
}
