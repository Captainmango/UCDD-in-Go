package usecases

import (
	"fmt"
	"testing"

	m "main/mocks"

	gomock "github.com/golang/mock/gomock"
)

func TestItCanHandleTheGetUserByIdUseCase(t *testing.T) {
    mockUserData := m.MakeMockUsersData()
    ctrl := gomock.NewController(t)

    mockService := m.NewMockIUsersService(ctrl)
    defer ctrl.Finish()

    mockService.EXPECT().
        GetUserById(int32(1)).
        Return(mockUserData[0])

    usecase := CreateGetUserByIdUseCase(mockService)
    _, err := usecase.Handle(int(1))

    if err != nil {
        t.Errorf("Got the following error: %v", err)
    } else {
        fmt.Println("GetUserByIdUseCase@Handle worked as expected")
    }
}