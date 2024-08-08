package bmdb

import (
	"context"
	"gorm.io/gorm"
)

const (
	defaultDbKey  dbContextKey = "defaultDb"
	businessDbKey dbContextKey = "businessDb"
)

// WithDefaultDbContext will inject the default *gorm.DB into the context.Context
// @param ctx context.Context
// @param name string the name of the MultiDbManager.RegisterDb name
// @return context.Context
// @return error
func WithDefaultDbContext(ctx context.Context, name string) (context.Context, error) {
	db, err := DefaultDbManager.Get(name)
	if err != nil {
		return ctx, err
	}
	return WithDbContext(ctx, defaultDbKey, db), nil
}

// MustWithDefaultDbContext will inject the default *gorm.DB into the context.Context, panics if not found
// @param ctx context.Context
// @param name string the name of the MultiDbManager.RegisterDb name
// @return context.Context
func MustWithDefaultDbContext(ctx context.Context, name string) context.Context {
	ctx, err := WithDefaultDbContext(ctx, name)
	if err != nil {
		panic(err)
	}
	return ctx
}

// GetDefaultDbFromContext will retrieve the default *gorm.DB from the context.Context
func GetDefaultDbFromContext(ctx context.Context) (*gorm.DB, error) {
	db, err := GetDbFromContext(ctx, defaultDbKey)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// MustGetDefaultDbFromContext will retrieve the default *gorm.DB from the context.Context, panics if not found
// @param ctx context.Context
// @return *gorm.DB
func MustGetDefaultDbFromContext(ctx context.Context) *gorm.DB {
	db, err := GetDefaultDbFromContext(ctx)
	if err != nil {
		panic(err)
	}
	return db
}

// WithBusinessDbContext will inject the business *gorm.DB into the context.Context
func WithBusinessDbContext(ctx context.Context, name string) (context.Context, error) {
	db, err := DefaultDbManager.Get(name)
	if err != nil {
		return ctx, err
	}
	return WithDbContext(ctx, businessDbKey, db), nil
}

// MustWithBusinessDbContext will inject the business *gorm.DB into the context.Context, panics if not found
func MustWithBusinessDbContext(ctx context.Context, name string) context.Context {
	ctx, err := WithBusinessDbContext(ctx, name)
	if err != nil {
		panic(err)
	}
	return ctx
}

// GetBusinessDbFromContext from context get business db
func GetBusinessDbFromContext(ctx context.Context) (*gorm.DB, error) {
	return GetDbFromContext(ctx, businessDbKey)
}

// MustGetBusinessDbFromContext from context get business db
func MustGetBusinessDbFromContext(ctx context.Context) *gorm.DB {
	db, err := GetBusinessDbFromContext(ctx)
	if err != nil {
		panic(err)
	}
	return db
}

// MustWithTenantDbContext inject default db and business db into context
// @param ctx context.Context
// @param defaultName string of the MultiDbManager.RegisterDb name
// @param businessName string of the MultiDbManager.RegisterDb name
// @return context.Context
func MustWithTenantDbContext(ctx context.Context, defaultName, businessName string) context.Context {
	ctx = MustWithDefaultDbContext(ctx, defaultName)
	ctx = MustWithBusinessDbContext(ctx, businessName)
	return ctx
}

// MustGetTenantDbFromContext from context get default db and business db
// @param ctx context.Context
// @return *gorm.DB defaultDb and businessDb
func MustGetTenantDbFromContext(ctx context.Context) (defaultDb *gorm.DB, businessDb *gorm.DB) {
	defaultDb = MustGetDefaultDbFromContext(ctx)
	businessDb = MustGetBusinessDbFromContext(ctx)
	return
}
