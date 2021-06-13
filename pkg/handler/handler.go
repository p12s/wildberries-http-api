package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/p12s/wildberries-http-api/docs"
	"github.com/p12s/wildberries-http-api/pkg/service"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func limitMaxConnections(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}

func (h *Handler) InitRoutes(limitMaxConnection int) *gin.Engine {
	router := gin.New()
	router.Use(limitMaxConnections(limitMaxConnection)) // уменьшение максимального кол-ва может ускорить ответ. Но надо тестировать

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
