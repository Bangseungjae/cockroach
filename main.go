package main

import (
	"Bangseungjae/cockroach/config"
	"Bangseungjae/cockroach/database"
	"Bangseungjae/cockroach/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
