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

	dns := `user="scenter01" 
	password="scenter01new" 
	connectString="10.2.1.98:1571/RUBRAM?expire_timconnect_time=2"
	sysdba=0
	sysoper=0
	poolMinSessions=1
	poolMaxSessions=1000
	poolIncrement=1
	standaloneConnection=0
	enableEvents=0
	heterogeneousPool=0
	externalAuth=0
	prelim=0
	poolWaitTimeout=5m
	poolSessionMaxLifetime=1h
	poolSessionTimeout=30s
	timezone="local"`
	
	driver := "godror"

	return sqlx.Open(driver, dns)

}
