package bmdb

import (
	"gorm.io/gorm"
)

// DbManager is a manager for multiple db instances
type DbManager interface {
	// Register a db with name and sign
	// @param name: db name
	// @param sign: db connection config sign, used to re-register db when connection config changed
	// @param db: gorm db instance
	Register(name string, sign string, db *gorm.DB)
	// Unregister remove db by name
	// @param name: db name
	// @return: true if db exists and removed, false if db not exists
	Unregister(name string) bool
	// Get get db by name
	// @param name: db name
	// @return: db instance and nil if db exists, nil and error if db not exists
	Get(name string) (*gorm.DB, error)
	// GetSign get db sign by name
	// @param name: db name
	// @return: db sign and nil if db exists, nil and error if db not exists
	GetSign(name string) (string, error)
	// Exists check if db exists by name
	// @param name: db name
	// @return: true if db exists, false if db not exists
	Exists(name string) bool
	// Count get db count
	Count() int
}
