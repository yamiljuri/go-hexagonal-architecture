package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yamiljuri/hex_arq/internal/core/ports"
	"github.com/yamiljuri/hex_arq/internal/core/services/personssrv"
	"github.com/yamiljuri/hex_arq/internal/core/services/usersrv"
	"github.com/yamiljuri/hex_arq/internal/infrastructure/repositories/memory"
	"github.com/yamiljuri/hex_arq/internal/infrastructure/repositories/sql"
)


func RegisterRouter(engine *gin.Engine){

	repoPerson := memory.NewPersonsRepository()
	servPerson := personssrv.NewService(repoPerson)
	getEndpoint := GetPersonEndpoint(servPerson)

	engine.GET("person/:id", getEndpoint)

}