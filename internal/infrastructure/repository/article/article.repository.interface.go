package article

import (
	"context"
	articleDomain "todo/internal/article/domain"
)

type ArticleRepository interface {
	GetAll(ctx context.Context, page int64, limit int64) (articles []*articleDomain.Article, total int64, err error)
	GetById(ctx context.Context, id int64) (article *articleDomain.Article, err error)
}
