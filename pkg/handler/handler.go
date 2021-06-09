package handler

import (
	"github.com/p12s/wildberries-http-api/pkg/service"
	"github.com/gin-gonic/gin"
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

	v1 := router.Group("/api")
	{
		api := router.Group("/v1", nil) //, h.userIdentity
		{
			auth := router.Group("/auth")
			{
				auth.POST("/sign-up", h.signUp)
				auth.POST("/sign-in", h.signIn)
			}

			user := api.Group("/user")
			{
				user.POST("/", h.createUser)      // добавление
				user.GET("/", h.getAllUsers)      // просмотр всех
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
	}
	//v1.GET("/hello", h.test)

	_ = v1
	return router
}
