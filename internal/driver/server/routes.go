package driverServer

import (
	articleDelivery "todo/internal/article/delivery"
	articleRepository "todo/internal/article/infras"
	articleServices "todo/internal/article/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (server *Server) MapRoutes(echo *echo.Echo) error {
	// Load repos
	articleRepo := articleRepository.NewArticleRepository(server.db)

	// Load services
	articleService := articleServices.NewArticleService(articleRepo)

	// Load handlers
	articleHandler := articleDelivery.NewArticleHandler(articleService)

	// Load middlewares
	echo.Use(middleware.BodyLimit("2M"))

	// Versioning
	v1 := echo.Group("/api/v1")

	// Load group routes
	articleGroup := v1.Group("/articles")

	articleDelivery.ArticleRoutes(articleGroup, articleHandler)

	return nil
}
