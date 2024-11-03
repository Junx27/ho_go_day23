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
	filter := c.Query("filter")

	if filter != "" {
		query = fmt.Sprintf("%s WHERE user_name = ?", query)
	}

	if filter != "" {
		err := model.DB.Raw(query, filter).Scan(&users).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.BaseResponse{
				Message: fmt.Sprintf("failed to get users: %s", err.Error()),
			})
			return
		}
	} else {
		err := model.DB.Raw(query).Scan(&users).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.BaseResponse{
				Message: fmt.Sprintf("failed to get users: %s", err.Error()),
			})
			return
		}
	}

	c.JSON(http.StatusOK, model.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    users,
	})
}

func CreateUserHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BaseResponse{
			Message: fmt.Sprintf("failed to bind request: %s", err.Error()),
		})
		return
	}

	err = model.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.BaseResponse{
			Message: fmt.Sprintf("failed to save user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, model.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    user,
	})
}
