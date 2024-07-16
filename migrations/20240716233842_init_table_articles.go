package migrations

import (
	articleDomain "todo/internal/article/domain"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "20240716233842_init_table_articles",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&articleDomain.Article{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&articleDomain.Article{})
		},
	})
}
