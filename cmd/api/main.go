package main

import (
	"golang-restaurant-management/internal/config"
	"golang-restaurant-management/internal/database"
	"golang-restaurant-management/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var cfg *config.Config

	var err error
	cfg, err = config.Load()

	if err != nil {
		log.Println("failed to load configuration: ", err)
		return
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Println("failed to connect to database: ", err)
		return
	}
	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Restaurant management API is running",
			"status":   "success",
			"database": "connected",
		})
	})

	routes.UserRoutes(router, pool, cfg)

	routes.FoodRoutes(router, pool)
	routes.MenuRoutes(router, pool)
	routes.TableRoutes(router, pool)
	routes.OrderRoutes(router, pool)
	routes.OrderItemRoutes(router, pool)
	routes.InvoiceRoutes(router, pool)

	router.Run(":" + cfg.Port)
}
