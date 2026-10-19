package routes

import (
	"golang-restaurant-management/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func OrderItemRoutes(router *gin.Engine, pool *pgxpool.Pool) {
	router.GET("/orderItems", controllers.GetOrderItems(pool))
	router.GET("/orderItems/:orderItem_id", controllers.GetOrderItem(pool))
	router.POST("/orderItems/:orderItem_id", controllers.CreateOrderItem(pool))
	router.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem(pool))
}
