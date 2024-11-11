package db

import "fmt"

func NewConfig() *DBConfig {
	return &DBConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "hao",
		Password: "020899",
		DBName:   "gai-scraper",
	}
}

func (c *DBConfig) BuildConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}
