package main

import (
	"github.com/fahad-md-kamal/go-bitly/models"
	"github.com/fahad-md-kamal/go-bitly/server"
)

func main(){
	models.Setup()
	server.SetupAndListen()
}