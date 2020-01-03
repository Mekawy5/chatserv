package application

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ApplicationService struct {
	Repository *ApplicationRepository
}

func NewApplicationService(r *ApplicationRepository) *ApplicationService {
	return &ApplicationService{
		Repository: r,
	}
}

func (s *ApplicationService) GetAll() []ApplicationModel {
	return s.Repository.GetAll()
}

func (s *ApplicationService) Get(id uint) ApplicationModel {
	return s.Repository.Get(id)
}

func (s *ApplicationService) Save(app ApplicationModel) ApplicationModel {
	return s.Repository.Save(app)
}
