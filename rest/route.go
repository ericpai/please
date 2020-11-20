package rest

import (
	"net/http"

	"github.com/ericpai/please/rest/controller"
	"github.com/ericpai/please/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(taskController *controller.TaskController) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api", middleware.Request, middleware.Response)
	api.GET("/tasks", controller.Wrapper(taskController.GetAll))
	api.GET("/tasks/:task_id", controller.Wrapper(taskController.GetByID))
	api.POST("/tasks", controller.Wrapper(taskController.Create))
	api.PATCH("/tasks/:task_id", controller.Wrapper(taskController.Update))
	api.DELETE("/tasks/:task_id", controller.Wrapper(taskController.Delete))
	router.StaticFS("/assets", http.Dir("./assets"))
	router.StaticFile("/", "index.html")
	return router
}
