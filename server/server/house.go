package server

import (
	"log"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startHouse(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/house", func(c *gin.Context) {
		var queryInfo QueryHouse
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindHouses(dbc, queryInfo.BuildNo,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/house/add", func(c *gin.Context) {
		var add table.House
		err := c.BindJSON(&add)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			panic(err)
		} else {
			c.JSON(http.StatusOK, table.InsertHouse(dbc, add))
		}
	})
	router.POST("/house/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindHouseKVs(dbc, params))
		}
	})
}
