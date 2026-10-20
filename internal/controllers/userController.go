package controllers

import (
	"golang-restaurant-management/internal/models"
	"golang-restaurant-management/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"email"`
	Phone      string `json:"phone" binding:"phone"`
}
type LoginRequest struct {
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"email"`
	Phone      string `json:"phone" binding:"phone"`
}

func GetUsers(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetUser(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerRequest RegisterRequest
		if err := c.BindJSON(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
			return
		}
		if len(registerRequest.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 characters long"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hashed password" + err.Error()})
			return
		}
		user := &models.User{
			First_name: registerRequest.First_name,
			Last_name:  registerRequest.Last_name,
			Password:   string(hashedPassword),
			Email:      registerRequest.Email,
			Phone:      registerRequest.Phone,
		}
		createdUser, err := repository.CreateUser(pool, user)
		if err != nil {
			if err.Error() != "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Email already has registered"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, createdUser)
	}
}
func Login(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
