package utils

import (
	"fmt"
	"github.com/thifnmi/go-book-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type MysqlStorage struct {
	config *config.AppConfig
	db     *gorm.DB
}

func NewMysqlStorage(appConfig *config.AppConfig) *MysqlStorage {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       appConfig.MySQLURI, // data source name
		DefaultStringSize:         256,                // default size for string fields
		DisableDatetimePrecision:  true,               // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,               // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		fmt.Sprintf("Error when connect to mysql server %v", err)
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	instance := &MysqlStorage{
		config: appConfig,
		db:     db,
	}

	return instance
}

func (m *MysqlStorage) Config() *config.AppConfig {
	return m.config
}

func (m *MysqlStorage) Db() *gorm.DB {
	return m.db
}

func (m *MysqlStorage) Close() {
	db, _ := m.db.DB()
	db.Close()
}

func (m *MysqlStorage) Ping() {
	db, _ := m.db.DB()
	err := db.Ping()
	if err != nil {
		fmt.Sprintf("Ping database failded with error %v", err)
	}

}
