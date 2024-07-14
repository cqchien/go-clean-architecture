package articleInterfaces

import (
	"context"
	articleDomain "todo/internal/article/domain"
)

type ArticleService interface {
	GetAll(ctx context.Context, page int, limit int) (articles []*articleDomain.Article, total int64, err error)
}
