package settings

import (
	"avd.com/todo/template_func_map"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Init(server *gin.Engine) {
	server.SetFuncMap(template.FuncMap{
		"dateFormat": template_func_map.DateFormat,
	})

	server.LoadHTMLGlob("templates/**/*")
	server.Static("/static", "static")
}
