package core

import (
	"database/sql"
	"github.com/go-redis/redis"
	db "github.com/wecanooo/kora/app/models"
)

type DBConn struct {
	*sql.DB
}

// NewDBConn database instance 생성
func NewDBConn(conn *sql.DB) {
	defaultConnection = &DBConn{conn}
}

// NewStore store interface 생성
func NewStore(st db.Store) {
	store = st
}

// NewRedis redis instance 생성
func NewRedis(r *redis.Client) {
	redisClient = r
}
