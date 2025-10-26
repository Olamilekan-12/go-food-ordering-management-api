package routes

import (
	controller "github.com/Olamilekan-12/go-food-ordering-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.Getorder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
