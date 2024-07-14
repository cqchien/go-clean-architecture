package articleInterfaces

import (
	"context"
	articleDomain "todo/internal/article/domain"
)

type ArticleRepository interface {
	GetAll(ctx context.Context, page int, limit int) (articles []* articleDomain.Article, total int64, err error)
	// GetById(ctx context.Context, id int64) (article *articleDomain.Article, err error)
}
