package services

import (
	"gorm.io/gorm"

	"simple-api-gorm/config"
	"simple-api-gorm/models"
)

type Users struct {
}

func (s *Users) ById(id int) (*gorm.DB, models.User, error) {
	var model models.User
	result := config.Db().First(&model, id)
	return result, model, nil
}

func (s *Users) All(user interface{}) (*gorm.DB, []*models.User) {
	var list []*models.User
	result := config.Db().Where(user).Find(&list)
	return result, list
}

func (s *Users) Create(user *models.User) (*gorm.DB, *models.User) {
	result := config.Db().Create(&user)
	return result, user
}

func (s *Users) Update(user *models.User) (*gorm.DB, *models.User) {
	result := config.Db().Save(&user)
	return result, user
}

func (s *Users) Delete(id int) (*gorm.DB, models.User) {
	var model models.User
	config.Db().First(&model, id)
	result := config.Db().Delete(&model)
	return result, model
}
