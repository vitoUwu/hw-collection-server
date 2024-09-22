package db

import (
	"hwc/env"

	"github.com/lucsky/cuid"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Gorm *gorm.DB
}

func NewDatabase() *Database {
	connectionString := env.GetDatabaseURL()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Car{})
	db.Exec(`
		CREATE OR REPLACE VIEW searches AS
			SELECT
				cars.id,
				cars.model_name,
				cars.code,
				cars.series_name,
				cars.series_code,
				cars.notes,
				cars.col_code
			FROM cars;

		CREATE INDEX IF NOT EXISTS index_cars_id ON cars USING gin(to_tsvector('simple', id));
		CREATE INDEX IF NOT EXISTS index_cars_model_name ON cars USING gin(to_tsvector('english', model_name));
		CREATE INDEX IF NOT EXISTS index_cars_code ON cars USING gin(to_tsvector('simple', code));
		CREATE INDEX IF NOT EXISTS index_cars_series_name ON cars USING gin(to_tsvector('english', series_name));
		CREATE INDEX IF NOT EXISTS index_cars_series_code ON cars USING gin(to_tsvector('simple', series_code));
		CREATE INDEX IF NOT EXISTS index_cars_notes ON cars USING gin(to_tsvector('english', notes));
		CREATE INDEX IF NOT EXISTS index_cars_col_code ON cars USING gin(to_tsvector('simple', col_code));

		CREATE EXTENSION IF NOT EXISTS pg_trgm;

		CREATE INDEX IF NOT EXISTS index_trgm_model_name ON cars USING gin (model_name gin_trgm_ops);
		CREATE INDEX IF NOT EXISTS index_trgm_code ON cars USING gin (code gin_trgm_ops);
		CREATE INDEX IF NOT EXISTS index_trgm_series_name ON cars USING gin (series_name gin_trgm_ops);
		CREATE INDEX IF NOT EXISTS index_trgm_series_code ON cars USING gin (series_code gin_trgm_ops);
		CREATE INDEX IF NOT EXISTS index_trgm_notes ON cars USING gin (notes gin_trgm_ops);
		CREATE INDEX IF NOT EXISTS index_trgm_col_code ON cars USING gin (col_code gin_trgm_ops);
	`)

	return &Database{
		Gorm: db,
	}
}

var Db = NewDatabase()

func (db *Database) GenerateId() string {
	return cuid.New()
}
