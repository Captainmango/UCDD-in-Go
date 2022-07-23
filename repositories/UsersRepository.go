package repositories

import (
	"errors"
	d "main/domain"
)

var (
    users = []d.User{
        {Id: 1, Name: "John"},
        {Id: 2, Name: "Tony"},
        {Id: 3, Name: "Fred"},
        {Id: 4, Name: "Angela"},
    }
)

type UsersRepository struct {
}

type IUsersRepository interface {
    GetAllUsers() []d.User
    GetUserById(id int32) (d.User, error)
}

func GetUsersRepositoryInstance() *UsersRepository{
    return &UsersRepository{}
}

func(ur *UsersRepository) GetAllUsers() []d.User {
    return users
}

func (ur *UsersRepository) GetUserById(id int32) (d.User, error) {
    for i := range users {
        if users[i].Id == id {
            return users[i], nil
        }
    }

    return d.User{}, errors.New("No user found with this ID")
}

