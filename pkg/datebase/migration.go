package database

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

type HasTableName interface {
	TableName() string
}

type HasIndexes interface {
	Indexes() map[string][]string
}

type HasUniqueIndexes interface {
	UniqueIndexes() map[string][]string
}

type HasForeignKeys interface {
	ForeignKeys() map[string]string
}

func MigrateTables(tx *gorm.DB, tables ...HasTableName) *gorm.DB {
	for _, t := range tables {
		if tx = tx.AutoMigrate(t); tx.Error != nil {
			return tx
		}
		if hi, ok := t.(HasIndexes); ok {
			for indexName, columns := range hi.Indexes() {
				if tx = tx.Model(t).AddIndex(indexName, columns...); tx.Error != nil {
					return tx
				}
			}
		}
		if hi, ok := t.(HasUniqueIndexes); ok {
			for indexName, columns := range hi.UniqueIndexes() {
				if tx = tx.Model(t).AddUniqueIndex(indexName, columns...); tx.Error != nil {
					return tx
				}
			}
		}
	}
	for _, t := range tables {
		if hf, ok := t.(HasForeignKeys); ok {
			for col, foreignKey := range hf.ForeignKeys() {
				if tx = tx.Model(t).AddForeignKey(col, foreignKey, "RESTRICT", "CASCADE"); tx.Error != nil {
					return tx
				}
			}
		}
	}
	return tx
}

type Migration struct {
	Options    *gormigrate.Options
	InitSchema gormigrate.InitSchemaFunc
	Migrations []*gormigrate.Migration
}

func (m *Migration) Migrate(db *gorm.DB) error {
	initialization := gormigrate.New(db, m.Options, m.Migrations)
	initialization.InitSchema(m.InitSchema)
	return initialization.Migrate()
}
