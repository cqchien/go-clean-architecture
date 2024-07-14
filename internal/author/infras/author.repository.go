package repository

import (
	"context"
	authorDomain "todo/internal/author/domain"

	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{db: db}
}

func (authorRepo *authorRepository) GetById(ctx context.Context, id int64) (author *authorDomain.Author, err error) {
	var authorRes authorDomain.Author
	errQueryAuthor := authorRepo.db.WithContext(ctx).Where(&authorDomain.Author{
		ID: id,
	}).First(&authorRes).Error

	if errQueryAuthor != nil {
		return nil, errQueryAuthor
	}

	return &authorRes, errQueryAuthor
}
