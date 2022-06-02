package repository

import (
	"crud/domain"
	"crud/entity"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(post entity.Post) error {
	r.DB.Create(&post)

	return nil
}

func (r *repository) FindAll() ([]entity.Post, error) {
	var post []entity.Post
	r.DB.Find(&post)

	return post, nil
}

func (r *repository) FindOne(id int) (entity.Post, error) {
	var posts entity.Post
	r.DB.Where("id = ?", id).First(&posts)

	return posts, nil
}

func (r *repository) Update(id int, post entity.Post) (entity.Post, error) {
	r.DB.Where("id = ?", id).Updates(&post)

	return post, nil
}

func (r *repository) Delete(id int, post entity.Post) error {
	r.DB.Where("id = ?", id).Delete(&post)

	return nil
}
