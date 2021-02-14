package core

import (
	"database/sql"
	"github.com/go-redis/redis"
)

type DBConn struct {
	*sql.DB
}

// NewDBConn database instance 생성
func NewDBConn(conn *sql.DB) {
	defaultConnection = &DBConn{conn}
}

// NewRedis redis instance 생성
func NewRedis(r *redis.Client) {
	redisClient = r
}
