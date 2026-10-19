package routes

import (
	"golang-restaurant-management/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MenuRoutes(router *gin.Engine, pool *pgxpool.Pool) {
	router.GET("/menus", controllers.GetMenus(pool))
	router.GET("/menus/:menu_id", controllers.GetMenu(pool))
	router.POST("/menus/:menu_id", controllers.CreateMenu(pool))
	router.PATCH("/menus/:menu_id", controllers.UpdateMenu(pool))
}
