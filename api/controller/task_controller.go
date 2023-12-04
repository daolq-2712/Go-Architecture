package controller

import (
	"Go-Architecture/common"
	"Go-Architecture/domain"
	"Go-Architecture/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc TaskController) Create(c *gin.Context) {
	// Bind task
	var task entity.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return
	}

	// Get UserId
	userID := c.GetString("x-user-id")
	task.UserID = userID

	// Create task
	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (tc TaskController) Fetch(c *gin.Context) {
	// TODO implement later
}
