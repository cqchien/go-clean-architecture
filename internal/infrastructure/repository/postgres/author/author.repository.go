package author

import (
	"context"

	"gorm.io/gorm"

	domainAuthor "todo/internal/author/domain"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{db: db}
}

func (authorRepo *authorRepository) GetById(ctx context.Context, id int64) (author domainAuthor.Author, err error) {
	var authorRes domainAuthor.Author
	errQueryAuthor := authorRepo.db.WithContext(ctx).Where(&domainAuthor.Author{
		ID: id,
	}).First(&authorRes).Error

	// if errQueryAuthor != nil {
	// 	return nil, err
	// }

	return authorRes, errQueryAuthor
}
