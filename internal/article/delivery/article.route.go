package articleDelivery

import (
	articleInterfaces "todo/internal/article/interfaces"

	"github.com/labstack/echo/v4"
)

func ArticleRoutes(group *echo.Group, articleHandler articleInterfaces.ArticleHandlerInterface) {
	group.GET("", articleHandler.GetAll())
}
