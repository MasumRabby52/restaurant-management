package routes

import (
	"golang-restaurant-management/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func FoodRoutes(router *gin.Engine, pool *pgxpool.Pool) {
	router.GET("/foods", controllers.GetFoods(pool))
	router.GET("/foods/:food_id", controllers.GetFood(pool))
	router.POST("/foods", controllers.CreateFood(pool))
	router.PATCH("/foods/:food_id", controllers.UpdateFood(pool))
}
