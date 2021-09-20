package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

var cache redis.Conn

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}

	cache = conn
}

func main()  {
	initCache()
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
