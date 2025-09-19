package service

import "blog/model"

// 获取文章内容
func FindPostsByCondition(post *model.Post) []model.Post {
	db := loadConfGetDb()
	var posts []model.Post
	tx := db.Preload("User")
	if post != nil {
		if (*post).ID != 0 {
			tx.Where("id = ?", (*post).ID)
		}
		if (*post).Title != "" {
			tx.Where("title like ?", "%"+(*post).Title+"%")
		}
		if (*post).Content != "" {
			tx.Where("content like ?", "%"+(*post).Content+"%")
		}
		if (*post).UserID != 0 {
			tx.Where("user_id = ?", (*post).UserID)
		}
	}
	tx.Find(&posts)
	return posts
}

// 更新文章
func UpdatePost(post *model.Post) error {
	db := loadConfGetDb()
	return db.Save(post).Error
}

// 删除文章
func DeletePost(id uint64) error {
	db := loadConfGetDb()
	return db.Delete(&model.Post{}, id).Error
}

// 创建文章
func CreatePost(post *model.Post) error {
	db := loadConfGetDb()
	return db.Create(post).Error
}

// 通过文章ID获取文章
func FindPostsById(id uint) model.Post {
	db := loadConfGetDb()
	var model model.Post
	db.First(&model, id)
	return model
}
