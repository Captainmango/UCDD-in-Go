package usecases

import (
	"encoding/json"
	d "main/domain"
	s "main/services"
)

type GetUserByIdUseCase struct {
    usersService s.IUsersService
}

type IGetUserByIdUseCase interface {
    Handle(userId int32) (string, error)
}

func (ubi *GetUserByIdUseCase) Handle(id int) (string, error) {
    user := ubi.usersService.GetUserById(int32(id))
    usersJson, err := serializeUserJson(user)

    if err != nil {
        return "", err
    }

    return usersJson, nil
}

func CreateGetUserByIdUseCase(us s.IUsersService) *GetUserByIdUseCase {
    return &GetUserByIdUseCase{us}
}

func serializeUserJson(data d.User) (string, error) {
    output, err := json.Marshal(data)

    if err != nil {
        return "", err
    }

    return string(output), nil
}