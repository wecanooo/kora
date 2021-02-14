package core

import (
	"database/sql"
	"github.com/go-redis/redis"
	db "github.com/wecanooo/kora/app/models"
)

type DBConn struct {
	*sql.DB
}

// NewDBConn creates and set a database instance
func NewDBConn(conn *sql.DB) {
	defaultConnection = &DBConn{conn}
}

// NewStore set a store interface
func NewStore(st db.Store) {
	store = st
}

// NewRedis set a redis instance
func NewRedis(r *redis.Client) {
	redisClient = r
}
