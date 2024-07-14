package router

import (
	"github.com/gin-gonic/gin"
	"server-farm/controler"
)

func WebSocket(incoming *gin.Engine) {
	incoming.GET("/:id/updateInventories", controler.InventoryUpdate())
}
