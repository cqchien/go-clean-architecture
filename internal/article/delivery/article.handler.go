package articleDelivery

import (
	"net/http"
	"strconv"
	articleInterfaces "todo/internal/article/interfaces"
	paginationConstant "todo/internal/constants"

	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	articleService articleInterfaces.ArticleService
}

func NewArticleHandler(articleService articleInterfaces.ArticleService) articleInterfaces.ArticleHandlerInterface {
	return &articleHandler{articleService: articleService}
}

// ResponseError represents the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// PaginationData represents paginated data response
type PaginationData struct {
	Data  any   `json:"data"`
	Total int64 `json:"total"`
}

func (articleHandler *articleHandler) GetAll() echo.HandlerFunc {
	return func(context echo.Context) error {
		page, errParsePage := strconv.Atoi(context.QueryParam("page"))
		limit, errParseLimit := strconv.Atoi(context.QueryParam("limit"))

		if errParsePage != nil {
			page = paginationConstant.PAGE_DEFAULT
		}

		if errParseLimit != nil {
			limit = paginationConstant.LIMIT_DEFAULT
		}

		articlesResponse, totalArticles, errQueryArticles := articleHandler.articleService.GetAll(context.Request().Context(), page, limit)

		if errQueryArticles != nil {
			return context.JSON(http.StatusInternalServerError, ResponseError{Message: errQueryArticles.Error()})
		}
		return context.JSON(http.StatusOK, PaginationData{Data: articlesResponse, Total: totalArticles})
	}
}
