package server

import (
	"log"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startEvent(router *gin.Engine, dbc *mgo.Database) {
	router.POST("/event", func(c *gin.Context) {
		var queryInfo QueryEvent
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindEvents(dbc, queryInfo.Index,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/event/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindEventKVs(dbc, params))
		}
	})
}
