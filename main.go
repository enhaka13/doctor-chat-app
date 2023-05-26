package main

import (
	"doctor-chat-app/auth"
	"doctor-chat-app/doctor"
	"doctor-chat-app/handler"
	"doctor-chat-app/helper"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "freedb_march:E6gRgU$Ca#aKCBn@tcp(sql.freedb.tech:3306)/freedb_doctorchat?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("database connected...")
	}

	doctorRepository := doctor.NewRepository(db)

	doctorService := doctor.NewService(doctorRepository)
	authService := auth.NewService()

	doctorHandler := handler.NewDoctorHandler(doctorService, authService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/register/doctors", doctorHandler.RegisterDoctor)

	router.Run()
}

func authMiddleware(authService auth.Service, doctorService doctor.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

		doctorID := int(claim["doctor_id"].(float64))

		doctor, err := doctorService.GetDoctorByID(doctorID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

		c.Set("currentUser", doctor)
	}
}
