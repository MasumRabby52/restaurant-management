package routes

import (
	"golang-restaurant-management/internal/config"
	"golang-restaurant-management/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UserRoutes(router *gin.Engine, pool *pgxpool.Pool, cfg *config.Config) {
	router.GET("/users", controllers.GetUsers(pool))
	router.GET("/users/:user_id", controllers.GetUser(pool))
	router.POST("/users/signup", controllers.SignUp(pool))
	router.POST("/users/login", controllers.Login(pool, cfg))
}
