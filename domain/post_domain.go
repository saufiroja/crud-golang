package domain

import "crud/entity"

type PostRepository interface {
	Create(post entity.Post) error
	FindAll() ([]entity.Post, error)
	FindOne(id int) (entity.Post, error)
	Update(id int, post entity.Post) (entity.Post, error)
	Delete(id int, post entity.Post) error
}

type PostService interface {
	Create(post entity.Post) error
	FindAll() (posts []entity.Post, err error)
	FindOne(id int) (entity.Post, error)
	Update(id int, post entity.Post) (entity.Post, error)
	Delete(id int, post entity.Post) error
}
