package service

import "github.com/nnanhkhoi/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserInfo() string {
	return us.userRepo.GetUserInfo()
}
