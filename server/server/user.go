package server

import (
	"fmt"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startUser(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/users", func(c *gin.Context) {
		var queryInfo QueryUser
		err := c.BindJSON(&queryInfo)
		fmt.Println("startUser", queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindUsers(dbc, queryInfo.UserName, queryInfo.Type,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})
}