package table

import (
	"log"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GovTableName 政府人员表格
const GovTableName = "Gov"

//Gov 政府人员信息
type Gov struct {
	UserName     string //用户名
	RealName     string //真实姓名
	Password     string //登录密码
	OfficeNumber string //工作电话
	Type         string //人员类型
}

//Govs 政府人员集
type Govs struct {
	Govs []Gov
}

//InsertGov 插入user
func InsertGov(db *mgo.Database, info Gov) string {
	c := db.C(GovTableName)
	count, err := c.Find(bson.M{"username": info.UserName}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return "用户名：" + info.UserName + "已存在, 请重新用户名"
	}
	err = c.Insert(&info)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindGovs 查询小区信息集
func FindGovs(db *mgo.Database, userName string, pageNo int, pageSize int) interface{} {
	c := db.C(GovTableName)
	var result []Gov
	var err error
	if userName == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"username": userName}).All(&result)
	}
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}