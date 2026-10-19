package routes

import (
	"golang-restaurant-management/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TableRoutes(router *gin.Engine, pool *pgxpool.Pool) {
	router.GET("/tables", controllers.GetTables(pool))
	router.GET("/tables/:table_id", controllers.GetTable(pool))
	router.POST("/tables/:table_id", controllers.CreateTable(pool))
	router.PATCH("/tables/:table_id", controllers.UpdateTable(pool))
}
