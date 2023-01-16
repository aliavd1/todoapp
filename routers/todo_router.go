package routers

import (
	"avd.com/todo/handlers"
	"github.com/gin-gonic/gin"
)

func TodoRouter(r *gin.RouterGroup) {
	todoHandler := handlers.TodoHandler{}

	r.GET("/", todoHandler.FindAll())

	r.GET("/:id", todoHandler.FindOne())

	r.POST("/", todoHandler.Create())

	r.PUT("/:id", todoHandler.Update())

	r.DELETE("/:id", todoHandler.Delete())
}
