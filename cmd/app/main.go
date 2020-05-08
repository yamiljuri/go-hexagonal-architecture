package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yamiljuri/hex_arq/internal/infrastructure/server"
	"log"
)

func main(){
	engine := gin.Default()
	server.RegisterRouter(engine)
	log.Fatal(engine.Run(":8080"))
}
