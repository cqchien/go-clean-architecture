package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "20240717214754_init_table_articles-hehe",
		Migrate: func(tx *gorm.DB) error {
			// Write your migration code here
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Write your rollback code here
			return nil
		},
	})
}
