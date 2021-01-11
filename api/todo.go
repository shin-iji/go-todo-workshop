package api

import (
	"fmt"
	"go-todo-workshop/database"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CreateTodo struct {
	Username string `json:"username" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

func GetAllLists(c *gin.Context) {
	var todoLists []database.Todo
	database.DB.Find(&todoLists)

	c.JSON(http.StatusOK, gin.H{"Result": todoLists})
}

func CreateTodoList(c *gin.Context) {
	var input CreateTodo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	todoList := database.Todo{Username: input.Username, Title: input.Title, Message: input.Message}
	database.DB.Create(&todoList)

	c.JSON(http.StatusOK, gin.H{"To Do List": todoList})
}

func GetTodoList(c *gin.Context) {
	var todoLists []database.Todo

	if err := database.DB.Where("username = ?", c.Query("username")).Find(&todoLists).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"List": todoLists})
}

func DeleteList(c *gin.Context) {
	var todoList database.Todo

	if err := database.DB.Where("id = ?", c.Param("id")).First(&todoList).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	database.DB.Delete(&todoList)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
