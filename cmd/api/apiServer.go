package main

import (
	"github.com/n0byk/marbu/app/httpApi"
	"github.com/n0byk/marbu/config"
	"github.com/n0byk/marbu/infra/database"
)

func init() {
	config.NewConfig()
	database.NewStorage()
}

func main() {
	httpApi.NewHttpApi()
}
