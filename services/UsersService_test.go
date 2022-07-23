package services

import (
	"fmt"
	"testing"

	m "main/mocks"
	gomock "github.com/golang/mock/gomock"
)

func TestItCanGetAllUsers(t *testing.T) {
    mockUserData := m.MakeMockUsersData()
    
    ctrl := gomock.NewController(t)
    mockRepository := m.NewMockIUsersRepository(ctrl)
    defer ctrl.Finish()

    mockRepository.EXPECT().
        GetAllUsers().
        Return(mockUserData).
        AnyTimes()

    usersService := GetUsersServiceInstance(mockRepository)

    res := usersService.GetAllUsers()

    if res[0] != mockUserData[0] {
        t.Errorf("Element one of the returned value does not match got %v wanted %v", 
                 res[0], 
                 mockUserData[0])
    } else {
        fmt.Println("It can get all users")
    }
}

func TestItCanGetAUserById(t *testing.T) {
    mockUserData := m.MakeMockUsersData()
    
    ctrl := gomock.NewController(t)
    mockRepository := m.NewMockIUsersRepository(ctrl)
    defer ctrl.Finish()

    mockRepository.EXPECT().
        GetUserById(gomock.Eq(int32(2))).
        Return(mockUserData[1], nil).
        AnyTimes()

    usersService := GetUsersServiceInstance(mockRepository)

    res := usersService.GetUserById(2)

    if res != mockUserData[1] {
        t.Errorf("User retrieved not correct. Expected %v, got %v", mockUserData[1], res)
    } else {
        fmt.Println("It can get a user by ID")
    }
}