package database

import "gorm.io/gorm"

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type DBO struct {
	db *gorm.DB
}
