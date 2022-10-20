package global

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

const Port = ":54432"

var DBCfg = "host=127.0.0.1 port=5435 user=umar password=blahblah1 dbname=blog sslmode=disable"

type Status struct {
	Msg string `json:"_status"`
}

var DB *sql.DB
var RDB *redis.Client

var Err error
