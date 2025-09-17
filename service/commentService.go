package service

import "blog/model"

// 创建评论
func CreateComment(comment *model.Comment) error {
	db := loadConfGetDb()
	err := db.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

// 获取某个文章下的所有评论吧
func FindCommentsByPostId(postId uint) []model.Comment {
	db := loadConfGetDb()
	var comments []model.Comment
	db.Preload("User").Preload("Post").Where("post_id = ?", postId).Find(&comments)
	return comments
}
