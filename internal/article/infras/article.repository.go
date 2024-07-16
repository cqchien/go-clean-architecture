package articleRepository

import (
	"context"
	articleDomain "todo/internal/article/domain"
	articleInterfaces "todo/internal/article/interfaces"

	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) articleInterfaces.ArticleRepository {
	return &articleRepository{db: db}
}

func (articleRepo *articleRepository) GetAll(ctx context.Context, page int, limit int) (articles []*articleDomain.Article, total int64, err error) {
	var articlesResponse []*articleDomain.Article
	var totalArticles int64

	offset := (page - 1) * limit

	errQueryArticles := articleRepo.db.Model(&articleDomain.Article{}).WithContext(ctx).Offset(offset).Limit(limit).Find(&articlesResponse).Error

	if errQueryArticles != nil {
		return nil, 0, errQueryArticles
	}

	errQueryCountArticles := articleRepo.db.WithContext(ctx).Model(&articleDomain.Article{}).Count(&totalArticles).Error

	if errQueryCountArticles != nil {
		return nil, 0, errQueryCountArticles
	}

	return articlesResponse, totalArticles, errQueryArticles
}
