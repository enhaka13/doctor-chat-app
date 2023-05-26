package handler

import (
	"doctor-chat-app/auth"
	"doctor-chat-app/doctor"
	"doctor-chat-app/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type doctorHandler struct {
	doctorService doctor.Service
	authService   auth.Service
}

func NewDoctorHandler(doctorService doctor.Service, authService auth.Service) *doctorHandler {
	return &doctorHandler{doctorService, authService}
}

func (h *doctorHandler) RegisterDoctor(c *gin.Context) {
	var input doctor.RegisterDoctorInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errrors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errrors}

		response := helper.APIResponse("Please complete all column in register form", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDoctor, err := h.doctorService.RegisterDoctor(input)
	if err != nil {
		response := helper.APIResponse("Register doctor account failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := h.authService.GenerateToken(newDoctor.ID)
	if err != nil {
		response := helper.APIResponse("Failed to generate token", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := doctor.FormatDoctor(newDoctor, token)

	response := helper.APIResponse("Doctor account has been registered", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)
}
