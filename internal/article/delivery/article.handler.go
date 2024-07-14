package articleDelivery

import (
	"net/http"
	articleDomain "todo/internal/article/domain"
	articleInterfaces "todo/internal/article/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type articleHandler struct {
	articleService articleInterfaces.ArticleService
}

func NewArticleHandler(articleService articleInterfaces.ArticleService) articleInterfaces.ArticleHandlerInterface {
	return &articleHandler{articleService: articleService}
}

func (articleHandler *articleHandler) GetAll() echo.HandlerFunc {
	return func(context echo.Context) error {
		page := context.QueryParam("page")
		limit := context.QueryParam("limit")

		logrus.Info(page, limit)
		return context.JSON(http.StatusOK, []articleDomain.Article{})
		// articlesResponse, totalArticles, errQueryArticles := articleHandler.articleService.GetAll(ctx, page, limit)
	}
}
