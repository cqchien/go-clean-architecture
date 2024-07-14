package delivery

import articleInterfaces "todo/internal/article/interfaces"

type articleHandler struct {
	articleService articleInterfaces.ArticleService
}
