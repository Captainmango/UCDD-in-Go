package services

import (
	d "main/domain"
	r "main/repositories"
)

type UsersService struct {
    UserRepository r.IUsersRepository
}

type IUsersService interface {
    GetAllUsers() []d.User
    GetUserById(id int32) d.User
}

func GetUsersServiceInstance(usersRepository r.IUsersRepository) *UsersService {
    return &UsersService{usersRepository}
}

func (us *UsersService) GetAllUsers() []d.User {
    return us.UserRepository.GetAllUsers()
}

func (us *UsersService) GetUserById(id int32) d.User {
    user, err := us.UserRepository.GetUserById(id)

    if err != nil {
        panic(err)
    }

    return user
}