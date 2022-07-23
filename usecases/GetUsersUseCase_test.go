package usecases

import (
	"fmt"
	"testing"

	m "main/mocks"

	gomock "github.com/golang/mock/gomock"
)

func TestItCanHandleTheGetAllUsersUseCase(t *testing.T) {
    mockUserData := m.MakeMockUsersData()
    ctrl := gomock.NewController(t)

    mockService := m.NewMockIUsersService(ctrl)
    defer ctrl.Finish()

    mockService.EXPECT().
        GetAllUsers().
        Return(mockUserData).
        AnyTimes()

    useCase := GetUsersUseCase{mockService}
    _, err := useCase.Handle()

    if err != nil {
        t.Errorf("Got the following error %v", err)
    } else {
        fmt.Println("GetUsersUseCase@Handle worked as expected")
    }
}