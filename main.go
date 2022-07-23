package main

import (
    r "main/repositories"
    s "main/services"
	uc "main/usecases"
	"net/http"
    "strconv"
	"github.com/gin-gonic/gin"
)

var (
    usersRepository = r.GetUsersRepositoryInstance()
    usersService = s.GetUsersServiceInstance(usersRepository)
    
    getUsers = uc.CreateGetUsersUseCase(usersService)
    getUserById = uc.CreateGetUserByIdUseCase(usersService)
)

func main() {
    
    r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
        users, err := getUsers.Handle()

        if err != nil {
            panic(err)
        }
        
		c.JSON(http.StatusOK, users)
	})

    r.GET("/users/:id", func(c *gin.Context) {
        userId, _ := strconv.Atoi(c.Param("id"))
        
        user, err := getUserById.Handle(userId)

        if err != nil {
            panic(err)
        }
        
		c.JSON(http.StatusOK, user)
	})
	r.Run()
}