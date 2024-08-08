package bmdb

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const (
	DriverTypeMysql    = "mysql"
	DriverTypeSqlite   = "sqlite"
	DriverTypePostgres = "postgres"
	DriverTypeMssql    = "mssql"
)

func DbOpen(driverType, dsn string) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch driverType {
	case DriverTypeMysql:
		dialector = mysql.Open(dsn)
	case DriverTypeSqlite:
		dialector = sqlite.Open(dsn)
	case DriverTypePostgres:
		// Correct the dialector for postgres
		dialector = mysql.Open(dsn)
	case DriverTypeMssql:
		dialector = sqlserver.Open(dsn)
	default:
		panic("unknown driver type")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

func MustDbOpen(driverType, dsn string) *gorm.DB {
	db, err := DbOpen(driverType, dsn)
	if err != nil {
		panic(err)
	}
	return db
}
