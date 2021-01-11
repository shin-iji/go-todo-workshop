package main

import (
	"go-todo-workshop/api"
	"go-todo-workshop/database"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	database.ConnectDatabase()

	r.GET("/", api.GetAllLists)
	r.GET("/user", api.GetTodoList)
	r.POST("/", api.CreateTodoList)
	r.DELETE("/user/:id", api.DeleteList)
	r.POST("/upload", api.Upload)
	r.StaticFS("/file", http.Dir("public"))

	r.Run()
}
