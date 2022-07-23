package mocks

import (
    d "main/domain"
)

func MakeMockUsersData() []d.User {
    return []d.User{
        {Id: 1, Name: "Test1"},
        {Id: 2, Name: "Test2"},
    }
}