package author

import (
	"context"
	domainAuthor "todo/internal/author/domain"
)

type AuthorRepository interface {
	GetById(ctx context.Context, id int64) (author *domainAuthor.Author, err error)
}
