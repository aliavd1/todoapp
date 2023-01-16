package main

import (
	"avd.com/todo/databases"
	"avd.com/todo/routers"
	"avd.com/todo/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("Server is starting...")
	server := gin.Default()

	fmt.Println("Settings is initializing...")
	settings.Init(server)

	fmt.Println("Database is initializing...")
	databases.Init()

	todoRouteGroup := server.Group("/todos")

	routers.TodoRouter(todoRouteGroup)

	if err := server.Run(":8000"); err != nil {
		log.Fatalln(err.Error())
	}
}
