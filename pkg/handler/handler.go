package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/wildberries-http-api/pkg/service"
	//"github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api/v1", h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.POST("/", h.createUser)      // добавление
			user.GET("/:id", h.getUserById)   // просмотр
			user.PUT("/:id", h.updateUser)    // редактирование
			user.DELETE("/:id", h.deleteUser) // удаление

			comments := user.Group(":id/comment")
			{
				comments.POST("/", h.createComment) // создание коммента
				comments.GET("/", h.getAllComments) // показ всех комментов
			}
		}

		comments := api.Group("comment")
		{
			comments.GET("/:id", h.getCommentById)   // получение коммента
			comments.PUT("/:id", h.updateComment)    // обновление
			comments.DELETE("/:id", h.deleteComment) // удаление
		}
	}

	return router
}
