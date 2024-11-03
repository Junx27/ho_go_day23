package router

import (
	"github.com/Junx27/ho_go_day23/service/usecase/user"
	"github.com/gin-gonic/gin"
)

func ServerRoutes() {
	r := gin.Default()
	users := r.Group("/users/")
	{
		users.POST("/", user.CreateUserHandler)
		users.GET("/", user.ReadUsersHandler)
	}
	_ = r.Run()
}
