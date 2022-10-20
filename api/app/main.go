package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"blog/internal/pkg/global"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

func main() {
	global.DB, global.Err = sql.Open("postgres", global.DBCfg)
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	router := http.NewServeMux()
	go router.HandleFunc("/api/blog", route)

	fmt.Println("Running blog service on " + global.Port)
	http.ListenAndServe(global.Port, router)
}
