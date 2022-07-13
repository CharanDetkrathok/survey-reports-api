package databaseConnection

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

type connection struct{}

func NewDatabaseConnection() *connection {

	return &connection{}

}

func (c *connection) RedisConnection() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		// Addr:     "redis:6379",
		Password: "admin-survey",
		DB:       0,

	})

}

func (c *connection) OracleConnection() (*sqlx.DB, error) {

	dns := `ใส่ data sourse name ตรงนี้`
	
	driver := "godror"

	return sqlx.Open(driver, dns)

}
