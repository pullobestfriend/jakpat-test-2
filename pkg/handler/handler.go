package handler

import (
	"jakpat-test-2/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecases *usecase.Usecase
}

func NewHandler(usecases *usecase.Usecase) *Handler {
	return &Handler{usecases: usecases}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authMiddleware := NewAuthMiddleware(*h.usecases)
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/signup", h.CreateUser)
			user.POST("/login", h.GetUserByNameAndPassword)
		}
		item := api.Group("/item", authMiddleware)
		{
			item.POST("/", h.AddItem)
			item.GET("/:id", h.GetItemByIdAndStatus)
			item.PUT("/:id", h.UpdateItemById)
			item.GET("/seller/:seller", h.GetItemsBySellerIdAndStatus)
			item.POST("/order/", h.CreateOrder)
			item.GET("/order/:id", h.GetOrderById)
			item.PUT("/order/:id", h.UpdateOrderById)
		}
	}

	return router
}
