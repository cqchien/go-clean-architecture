package articleServices

import (
	"context"
	articleDomain "todo/internal/article/domain"
	articleInterfaces "todo/internal/article/interfaces"
)

type articleService struct {
	articleRepository articleInterfaces.ArticleRepository
}

func NewArticleService(articleRepository articleInterfaces.ArticleRepository) articleInterfaces.ArticleService {
	return &articleService{articleRepository: articleRepository}
}

func (articleService *articleService) GetAll(ctx context.Context, page int, limit int) (articles []*articleDomain.Article, total int64, err error) {
	articlesResponse, totalArticles, errQueryArticles := articleService.articleRepository.GetAll(ctx, page, limit)

	return articlesResponse, totalArticles, errQueryArticles
}
