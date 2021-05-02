package controllers

import (
	"net/http"

	"to-do-api/internal/core/web"
	datamodels "to-do-api/internal/db/models"
	"to-do-api/internal/server/models"
)

type TodoController struct {
	web.Controller
}

func find(c *web.Controller, payload *interface{}, id string) error {
	return c.DbContext.Model(&datamodels.Todo{}).First(&payload, "id = ?", id).Error
}

func (ctrlr *TodoController) Get(c *web.Controller) {
	// swagger:route GET /users/{id} users getSingleUser
	//
	// get a user by userID
	//
	// This will show a user info
	//
	//     Responses:
	//       200: ApiResponse
	result := models.ApiResponse{StatusCode: http.StatusOK}

	id := c.Context.Param("id")
	if len(id) == 0 {
		if err := c.DbContext.Model(&datamodels.Todo{}).Find(&result.Payload).Error; err != nil {
			result.StatusCode = http.StatusInternalServerError
			result.Message = err.Error()
			c.Logger.Error(result.Message)
		}
	} else {
		if err := find(c, &result.Payload, c.Context.Param("id")); err != nil {
			result.StatusCode = http.StatusInternalServerError
			result.Message = err.Error()
			c.Logger.Error(result.Message)
		}
	}
	c.Context.JSON(result.StatusCode, result)
}

func (ctrlr *TodoController) Post(c *web.Controller) {
	result := models.ApiResponse{StatusCode: http.StatusOK}

	if err := c.Context.ShouldBindJSON(&result.Payload); err != nil {
		result.StatusCode = http.StatusInternalServerError
		result.Message = err.Error()
		c.Logger.Error(result.Message)
	} else {
		c.DbContext.Model(&datamodels.Todo{}).Create(&result.Payload)
	}
	c.Context.JSON(result.StatusCode, result)
}

func (ctrlr *TodoController) Put(c *web.Controller) {
	result := models.ApiResponse{StatusCode: http.StatusOK}

	if err := find(c, &result.Payload, c.Context.Param("id")); err != nil {
		result.StatusCode = http.StatusInternalServerError
		result.Message = err.Error()
	} else {
		c.DbContext.Model(&datamodels.Todo{}).Updates(&result.Payload)
	}
	c.Context.JSON(result.StatusCode, result)
}

func (ctrlr *TodoController) Delete(c *web.Controller) {
	result := models.ApiResponse{StatusCode: http.StatusOK}

	if err := find(c, &result.Payload, c.Context.Param("id")); err != nil {
		result.StatusCode = http.StatusInternalServerError
		result.Message = err.Error()
	} else {
		c.DbContext.Model(&datamodels.Todo{}).Delete(&result.Payload)
	}
	c.Context.JSON(result.StatusCode, result)
}
