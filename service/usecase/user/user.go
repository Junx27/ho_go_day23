package user

import (
	"fmt"
	"net/http"

	"github.com/Junx27/ho_go_day23/model"
	"github.com/gin-gonic/gin"
)

func ReadUsersHandler(c *gin.Context) {
	var users []model.User

	query := "SELECT * FROM users"
	var args []any
	filter := c.Query("filter")

	if filter != "" {
		query = fmt.Sprintf(
			"%s %s",
			query,
			"where user_name = ?",
		)

		args = append(args, filter)
	}

	err := model.DB.Raw(query, args...).Scan(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: fmt.Sprintf("failed to get users: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Success",
		Data:    users,
	})
}

func CreateUserHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: fmt.Sprintf("failed to bind request: %s", err.Error()),
		})
		return
	}

	err = model.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: fmt.Sprintf("failed to save user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Success",
		Data:    user,
	})
}
