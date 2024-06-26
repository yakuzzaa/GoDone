package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	docs := router.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.POST("/", h.createList)
			lists.PUT("/", h.updateList)
			lists.DELETE("/", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.POST("/", h.createItem)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}

		}
	}
	return router
}
