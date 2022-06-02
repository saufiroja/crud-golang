package service

import (
	"crud/config"
	"crud/domain"
	"crud/entity"
)

type service struct {
	conf config.Config
	r    domain.PostRepository
}

func NewPostService(conf config.Config, r domain.PostRepository) domain.PostService {
	return &service{
		conf: conf,
		r:    r,
	}
}

func (s *service) Create(post entity.Post) error {
	return s.r.Create(post)
}

func (s *service) FindAll() (posts []entity.Post, err error) {
	return s.r.FindAll()
}

func (s *service) FindOne(id int) (entity.Post, error) {
	return s.r.FindOne(id)
}

func (s *service) Update(id int, post entity.Post) (entity.Post, error) {
	return s.r.Update(id, post)
}

func (s *service) Delete(id int, post entity.Post) error {
	return s.r.Delete(id, post)
}
