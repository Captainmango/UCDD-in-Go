package usecases

import (
    d "main/domain"
    s "main/services"
    "encoding/json"
)

var (
    us s.IUsersService = s.GetUsersServiceInstance()
)

type GetUsersUseCase struct {
    UsersService s.IUsersService
}

type IGetUsersUseCase interface {
    Handle() (string, error)
}

func (uc *GetUsersUseCase) Handle() (string, error) {
    users := us.GetAllUsers()
    usersJson, err := serializeUsersJson(users)

    if err != nil {
        return "", err
    }

    return usersJson, nil
}

func CreateGetUsersUseCase() *GetUsersUseCase {
    return &GetUsersUseCase{us}
}

func serializeUsersJson(data []d.User) (string, error) {
    output, err := json.Marshal(data)

    if err != nil {
        return "", err
    }

    return string(output), nil
}