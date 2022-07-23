package usecases

import (
    d "main/domain"
    s "main/services"
    "encoding/json"
)

type GetUsersUseCase struct {
    usersService s.IUsersService
}

type IGetUsersUseCase interface {
    Handle() (string, error)
}

func (uc *GetUsersUseCase) Handle() (string, error) {
    users := uc.usersService.GetAllUsers()
    usersJson, err := serializeUsersJson(users)

    if err != nil {
        return "", err
    }

    return usersJson, nil
}

func CreateGetUsersUseCase(usersService s.IUsersService) *GetUsersUseCase {
    return &GetUsersUseCase{usersService}
}

func serializeUsersJson(data []d.User) (string, error) {
    output, err := json.Marshal(data)

    if err != nil {
        return "", err
    }

    return string(output), nil
}