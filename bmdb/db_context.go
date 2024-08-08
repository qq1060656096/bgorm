package bmdb

import (
	"context"
	"gorm.io/gorm"
)

type dbContextKey string

// WithDbContext injects *gorm.DB into context.Context
func WithDbContext(ctx context.Context, dbKey interface{}, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

// GetDbFromContext retrieves *gorm.DB from context.Context
func GetDbFromContext(ctx context.Context, dbKey interface{}) (*gorm.DB, error) {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		return nil, errDbNotFoundInContext
	}
	return db, nil
}

// MustGetDbFromContext retrieves *gorm.DB from context.Context, panics if not found
func MustGetDbFromContext(ctx context.Context, dbKey interface{}) *gorm.DB {
	db, err := GetDbFromContext(ctx, dbKey)
	if err != nil {
		panic(err)
	}
	return db
}
