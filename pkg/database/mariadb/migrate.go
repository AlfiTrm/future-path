package mariadb

import (
	"future-path/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Kepemilikan{},
		&entity.Role{},
		&entity.Sekolah{},
		&entity.Universitas{},
		&entity.FAQ{},
		&entity.Berita{},
		&entity.User{},
	); err != nil {
		return err
	}
	return nil
}
