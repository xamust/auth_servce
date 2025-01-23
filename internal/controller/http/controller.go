package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/xamops/auth/docs"
	"gitlab.com/xamops/auth/internal/controller/http/v1/handlers"
	"gitlab.com/xamops/auth/pkg/middleware/auth"
	"net/http"
)

func NewRouter(deps handlers.Dependencies) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	//router.Group("/api/v1/*", handlers.NewRoutes(deps)...)
	//router.Any("/api/v1/*w", gin.WrapH(http.StripPrefix("/api/v1", srv)))

	// api V1
	h := handlers.NewRoutes(deps)
	v1 := router.Group("/api/v1")
	v1.POST("/login", h.Login)
	v1.POST("/logout", auth.Auth(deps.AccessHandler), h.Logout)
	v1.POST("/refresh", auth.Refresh(deps.RefreshHandler), h.Refresh)
	v1.GET("/user/current", auth.Auth(deps.AccessHandler), h.GetCurrentUser)

	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	docs.SwaggerInfo.BasePath = "/api/v1"
	// todo @ check persist auth
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
