package articleInterfaces

import "github.com/labstack/echo/v4"

type ArticleHandlerInterface interface {
	GetAll() echo.HandlerFunc
}
