package handler

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/auth_v1"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/list_v1"
)

type ApiHandler struct {
	authClient auth_v1.AuthV1Client
	listClient list_v1.ListV1Client
	itemClient item_v1.ItemV1Client
	logger     *slog.Logger
}

func NewHandler(authClient auth_v1.AuthV1Client, listClient list_v1.ListV1Client, itemClient item_v1.ItemV1Client, logger *slog.Logger) *ApiHandler {
	return &ApiHandler{
		authClient: authClient,
		listClient: listClient,
		itemClient: itemClient,
		logger:     logger,
	}
}

func (h *ApiHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Разрешить все домены
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

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
	api.Use(h.authMiddleware())
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.POST("/", h.createList)
			lists.PUT("/", h.updateList)
			lists.DELETE("/", h.deleteList)
		}
		items := api.Group("/items")
		{
			items.GET("/:list_id", h.getAllItems)
			items.GET("/:list_id/item_id", h.getItemById)
			items.POST("/:list_id", h.createItem)
			items.PUT("/:list_id/:item_id", h.updateItem)
			items.DELETE("/:list_id/:item_id", h.deleteItem)
		}
	}
	return router
}
