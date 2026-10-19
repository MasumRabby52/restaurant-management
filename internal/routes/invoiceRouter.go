package routes

import (
	"golang-restaurant-management/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InvoiceRoutes(router *gin.Engine, pool *pgxpool.Pool) {
	router.GET("/invoices", controllers.GetInvoices(pool))
	router.GET("/invoices/:invoice_id", controllers.GetInvoice(pool))
	router.POST("/invoices/:invoice_id", controllers.CreateInvoice(pool))
	router.PATCH("/invoices/:invoices_id", controllers.UpdateInvoice(pool))
}
