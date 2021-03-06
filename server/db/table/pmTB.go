package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PMTableName table name
const PMTableName = "PM"

//PM 物业信息
type PM struct {
	ID                   bson.ObjectId `bson:"_id"`
	Name                 string        //物业公司名称
	StreetID             string        //所在街道
	CommunityID          string        //所在社区
	XQID                 string        //所在小区
	LegalPerson          string        //独立法人
	Tel                  string        //
	WuYeZiZhi            string        //物业资质
	WuYeXinZhi           string        //物业性质
	XQEnv                string        //小区环境
	XiaoQuCleaning       string        //小区保洁
	GreenVegetatio       string        //绿化植被
	GuanYangBaoHu        string        //管养保护
	XiaoFanJianCha       string        //消防检查
	DianTiBaoYang        string        //电梯保养
	CarParkInOrder       string        //车辆有序停放
	YeZhuCommunity       string        //业主委员会
	YeZhuCommunityTel    string        //业主委员会联系方式
	PastYearLevel        string        //上一年物业等级
	JianYiZhengGaiCuoShi string        //建议整改措施
	Img1                 string
	Img2                 string
	Img3                 string
	Img4                 string
	Img5                 string
}

//PMs 集合
type PMs struct {
	PMs []PM
}

//InsertPM 插入user
func InsertPM(db *mgo.Database, info PM) interface{} {
	c := db.C(PMTableName)
	count, err := c.Find(bson.M{"name": info.Name, "xqid": info.XQID}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "该Country已存在：" + info.Name + ", 请重新设置名称"}
	}
	info.ID = bson.NewObjectId()
	err = c.Insert(&info)
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindWuYes 查询小区信息集
func FindWuYes(db *mgo.Database, name string, pageNo int, pageSize int) interface{} {
	c := db.C(PMTableName)
	var result []PM
	var err error
	if name == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"name": name}).All(&result)
	}
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindWuYeKVs 通过一系列的K-V来查找记录
func FindWuYeKVs(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(PMTableName)
	var result []PM
	err := c.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if result == nil {
		return gin.H{"error": 1, "data": "没有查询到结果"}
	}
	return gin.H{"error": 0, "data": result}
}

//UpdatePM update street
func UpdatePM(db *mgo.Database, pm PM) interface{} {
	c := db.C(PMTableName)
	err := c.Update(bson.M{"_id": pm.ID}, pm)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//DelPMs 删除
func DelPMs(db *mgo.Database, ids []string) interface{} {
	c := db.C(PMTableName)
	var err error
	result := ""
	for _, v := range ids {
		err = c.Remove(bson.M{"_id": bson.ObjectIdHex(v)})
		if err != nil {
			log.Println(err.Error())
			result += err.Error()
			err = nil
		}
	}
	if result != "" {
		return gin.H{"error": 1, "data": result}
	}
	return gin.H{"error": 0, "data": Succ}
}
