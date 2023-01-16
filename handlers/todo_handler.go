package handlers

import (
	"avd.com/todo/schemas"
	"net/http"

	"avd.com/todo/databases"
	"avd.com/todo/models"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct{}

func (todo *TodoHandler) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var todoList []models.Todo
		var todoResponseList interface{}
		databases.DB.Find(&todoList)
		if todoList != nil {
			todoSchema := schemas.TodoSchema{
				Context: ctx,
				Object:  todoList,
			}
			todoResponseList = todoSchema.MakeResponse()
		} else {
			todoResponseList = []interface{}{}
		}
		ctx.HTML(http.StatusOK, "pages/home.tmpl", gin.H{
			"todoList": todoResponseList,
		})
	}
}

func (todo *TodoHandler) FindOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		todo := models.Todo{}
		databases.DB.Find(&todo, id)
		if todo.ID != 0 {
			todoSchema := schemas.TodoSchema{
				Context: ctx,
				Object:  todo,
			}
			todoResponse := todoSchema.MakeResponse()
			ctx.HTML(http.StatusOK, "pages/detail.tmpl", gin.H{
				"todo": todoResponse,
			})
		} else {
			ctx.Redirect(http.StatusOK, ctx.Request.Host+"/todos")
		}
	}
}

func (todo *TodoHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todo := models.Todo{}
		todoSchema := schemas.TodoSchema{
			Context: ctx,
			Object:  &todo,
		}
		todoSchema.MakeRequest()
		databases.DB.Create(&todo)
		todoResponse := todoSchema.MakeResponse()
		ctx.JSON(http.StatusCreated, gin.H{
			"todo": todoResponse,
		})
	}
}

func (todo *TodoHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		todo := models.Todo{}
		databases.DB.Find(&todo, id)
		if todo.ID != 0 {
			todoSchema := schemas.TodoSchema{
				Context: ctx,
				Object:  &todo,
			}
			todoSchema.MakeRequest()
			databases.DB.Save(&todo)
			ctx.JSON(http.StatusOK, gin.H{})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{})
		}
	}
}

func (todo *TodoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		todo := models.Todo{}
		databases.DB.Find(&todo, id)
		if todo.ID != 0 {
			databases.DB.Delete(&todo)
			ctx.JSON(http.StatusOK, gin.H{})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{})
		}
	}
}
