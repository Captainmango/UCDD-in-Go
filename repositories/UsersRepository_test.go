package repositories

import (
	"fmt"
	"testing"
)

var (
    usersRepository = GetUsersRepositoryInstance()
)

func TestItGetsAllUsers(t *testing.T) {
    res := usersRepository.GetAllUsers()

    if res != nil {
        fmt.Println("UsersRepository@GetAllUsers returns all users")
    } else {
        t.Error("Could not get users")
    }
}

func TestItCanGetUsersById(t *testing.T) {
    _, err := usersRepository.GetUserById(2)

    if err != nil {
        t.Errorf("Could not get user. Error: %v", err)
    } else {
        fmt.Println("UsersRepository@GetUserById returns a specific user by id")
    }
}

func TestItReturnsAnErrorIfTheUserIdDoesNotExist(t *testing.T) {
    _, err := usersRepository.GetUserById(-1)

    if err != nil {
        fmt.Println("UsersRepository@GetUserById returns an error message when ID does not exist")
    } else {
        t.Errorf("This did not throw an error when it should have")
    }
}