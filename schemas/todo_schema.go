package schemas

import (
	"avd.com/todo/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type TodoSchema struct {
	Context *gin.Context
	Object  interface{}
}

func (todoSchema *TodoSchema) MakeResponse() any {
	switch todoSchema.Object.(type) {
	case []models.Todo:
		var todoResponseList []todoResponse
		todoList := todoSchema.Object.([]models.Todo)
		for _, todo := range todoList {
			todoResponse := createResponse(todoSchema.Context, todo)
			todoResponseList = append(todoResponseList, *todoResponse)
		}
		return todoResponseList
	case models.Todo:
		todo := todoSchema.Object.(models.Todo)
		todoResponse := createResponse(todoSchema.Context, todo)
		return todoResponse
	case *models.Todo:
		todo := todoSchema.Object.(*models.Todo)
		todoResponse := createResponse(todoSchema.Context, *todo)
		return todoResponse
	default:
		return nil
	}
}

func (todoSchema *TodoSchema) MakeRequest() {
	todo := todoSchema.Object.(*models.Todo)
	todoReq := createRequest()
	err := todoSchema.Context.ShouldBindJSON(&todoReq)
	if err != nil {
		log.Fatalln("Binding json has errored")
	}
	createTodo(todo, todoReq)
}

type todoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type todoRequest struct {
	Title string `json:"title" form:"title"`
}

func createResponse(ctx *gin.Context, todo models.Todo) *todoResponse {
	return &todoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Url:       ctx.Request.URL.Path + strconv.Itoa(int(todo.ID)),
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func createRequest() *todoRequest {
	return &todoRequest{}
}

func createTodo(todo *models.Todo, todoReq *todoRequest) {
	todo.Title = todoReq.Title
}
