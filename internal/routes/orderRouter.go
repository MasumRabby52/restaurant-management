package routes

import (
	"golang-restaurant-management/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func OrderRoutes(router *gin.Engine, pool *pgxpool.Pool) {
	router.GET("/orders", controllers.GetOrders(pool))
	router.GET("/orders/:order_id", controllers.GetOrder(pool))
	router.POST("/orders/:order_id", controllers.CreateOrder(pool))
	router.PATCH("/orders/:order_id", controllers.UpdateOrder(pool))
}
