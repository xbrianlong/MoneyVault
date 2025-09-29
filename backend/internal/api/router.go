package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/health", healthCheckHandler)

	// // Grouping endpoints by resource
	// vaultGroup := r.Group("/api/transactions")
	// {
	// 	vaultGroup.GET("/", listTransactionHandler)             // list all transactions
	// 	vaultGroup.GET("/:id", viewTransactionHandler)          // view details of specific transaction
	// 	vaultGroup.POST("/", addTransactionHandler)             // add transaction manually
	// 	vaultGroup.POST("/uploadreceipt", uploadReceiptHandler) // batch adding transactions
	// 	vaultGroup.PUT("/:id", editTransactionHandler)          // edit specific transaction
	// 	vaultGroup.DELETE("/:id", deleteTransactionHandler)     // delete specific transaction
	// }

	// categoriesGroup := r.Group("/api/categories")
	// {
	// 	categoriesGroup.GET("/", listCategoriesHandler)       // list all categories
	// 	categoriesGroup.POST("/", addCategoryHandler)         // add new category
	// 	categoriesGroup.PUT("/:id", editCategoryHandler)      // edit specific category
	// 	categoriesGroup.DELETE("/:id", deleteCategoryHandler) // delete specific category
	// }

	// summaryGroup := r.Group("/api/summary")
	// {
	// 	summaryGroup.GET("/", getSummaryHandler) // get summary of transactions (include ability for params)
	// }

	return r
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
