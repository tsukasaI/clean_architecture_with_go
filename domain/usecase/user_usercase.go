package usecase

import (
	"clean_architecture_with_go/domain/entity"
	"clean_architecture_with_go/repository"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	GetUsers(ctx *gin.Context) ([]entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetUsers(ctx *gin.Context) ([]entity.User, error) {
	users, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
