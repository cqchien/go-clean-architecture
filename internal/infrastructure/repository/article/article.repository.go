package article

import (
	"context"
	articleDomain "todo/internal/article/domain"

	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db: db}
}

func (articleRepo *articleRepository) GetAll(ctx context.Context, page int64, limit int64) (articles []*articleDomain.Article, total int64, err error) {
	var articlesResponse []*articleDomain.Article
	var totalArticles int64

	offset := (page - 1) * limit

	errQueryArticles := articleRepo.db.WithContext(ctx).Offset(int(offset)).Limit(int(limit)).Find(&articlesResponse).Error

	if errQueryArticles != nil {
		return nil, 0, errQueryArticles
	}

	errQueryCountArticles := articleRepo.db.WithContext(ctx).Count(&totalArticles).Error

	if errQueryCountArticles != nil {
		return nil, 0, errQueryCountArticles
	}

	return articlesResponse, totalArticles, errQueryArticles
}
