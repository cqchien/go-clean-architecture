package authorInterfaces

import (
	"context"
	authorDomain "todo/internal/author/domain"
)

type AuthorRepository interface {
	GetById(ctx context.Context, id int64) (author *authorDomain.Author, err error)
}
