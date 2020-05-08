package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yamiljuri/hex_arq/internal/core/domain"
	"github.com/yamiljuri/hex_arq/internal/core/ports"
	"net/http"
	"strconv"
)

func GetPersonEndpoint(service ports.PersonsService) func(c *gin.Context){
	return func(c *gin.Context) {
		var person domain.Person
		id ,isExist := c.Params.Get("id")
		if isExist {
			idNumb, err := strconv.Atoi(id)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			person , err = service.Get(idNumb)
			if err != nil {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}
		c.JSON(http.StatusOK, BuildResponsePersonGet(person))
	}
}

//func GetPersonEndpoint(service ports.PersonsService) func(c *gin.Context){
//	return func(c *gin.Context) {
//		var person domain.Person
//		id ,isExist := c.Params.Get("id")
//		if isExist {
//			idNumb, err := strconv.Atoi(id)
//			if err != nil {
//				c.AbortWithStatus(http.StatusNotFound)
//				return
//			}
//			person , err = service.Get(idNumb)
//			if err != nil {
//				c.AbortWithStatus(http.StatusNoContent)
//				return
//			}
//		}
//		return person, nil
//	}
//}