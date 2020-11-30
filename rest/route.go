package rest

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ericpai/please/rest/controller"
	"github.com/ericpai/please/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(taskController *controller.TaskController) *gin.Engine {
	router := gin.Default()
	router.StaticFile("/", "web/build/index.html")
	if err := filepath.Walk("web/build", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			p := strings.TrimPrefix(path, "web/build")
			log.Printf("Register static file with path: %s[%s]\n", p, path)
			router.StaticFile(p, path)
		}
		return nil
	}); err != nil {
		panic(err.Error())
	}
	api := router.Group("/api", middleware.Request, middleware.Response)
	api.GET("/tasks", controller.Wrapper(taskController.GetAll))
	api.GET("/tasks/:task_id", controller.Wrapper(taskController.GetByID))
	api.POST("/tasks", controller.Wrapper(taskController.Create))
	api.PATCH("/tasks/:task_id", controller.Wrapper(taskController.Update))
	api.DELETE("/tasks/:task_id", controller.Wrapper(taskController.Delete))
	return router
}
