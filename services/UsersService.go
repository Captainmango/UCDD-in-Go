package services

import (
	d "main/domain"
	r "main/repositories"
)

var (
    usersRepository = r.GetUsersRepositoryInstance()
)

type UsersService struct {
    UserRepository r.IUsersRepository
}

type IUsersService interface {
    GetAllUsers() []d.User
    GetUserById(id int32) d.User
}

func (us *UsersService) GetAllUsers() []d.User {
    return us.UserRepository.GetAllUsers()
}

func GetUsersServiceInstance() *UsersService {
    return &UsersService{usersRepository}
}

func (us *UsersService) GetUserById(id int32) d.User {
    user, err := us.UserRepository.GetUserById(id)

    if err != nil {
        panic(err)
    }

    return user
}