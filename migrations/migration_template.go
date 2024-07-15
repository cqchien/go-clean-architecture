package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "TIMESTAMP_DESCRIPTION",
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
