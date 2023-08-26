package repository

import (
	"clean_architecture_with_go/domain/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)
	r.db.AutoMigrate(&entity.User{})
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
