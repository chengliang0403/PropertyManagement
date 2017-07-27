package table

import (
	"log"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//HouseTabelName table name
const HouseTabelName = "House"

//House 建筑信息表
type House struct {
	BuildNo             string //建筑编号
	Owner               string //房屋登记人
	Street              string //所属街道
	Community           string //所属社区
	XQ                  string //房屋所属小区名
	HouseBuildNo        string //房屋楼栋号
	HouseNo             string //房屋门牌号
	HouseType           string //房屋类型
	Year                string //建筑年代
	UseChange           string //使用变更
	MainCrack           string //主体结构裂缝
	FoundationDown      string //地基沉降变形
	MainSlant           string //主体结构倾斜
	CantileverCrack     string //悬挑结构破坏
	ParapetOff          string //女儿墙脱落
	OuterLloatedCoatOff string //外墙抹灰层剥落
	HouseDeform         string //房顶变形
	Disaster            string //地质灾害
	DisasterManage      string //地址灾害治理情况
	DrainageSsystem     string //排水系统
	InnerChange         string //房屋内部装修主体结构变更破坏情况
	IllegalBuild        string //违规搭建加层
	RankAppraisal       string //房屋鉴定等级
	Img1                string //图片
	Img2                string
	Img3                string
	Img4                string
	Img5                string //图片
	Img6                string
	Img7                string
	Img8                string
	Img9                string
}

//Houses 集合
type Houses struct {
	Houses []House
}

//InsertHouseInfoTB 插入
func InsertHouseInfoTB(db *mgo.Database, house House) string {
	c := db.C(HouseTabelName)
	count, err := c.Find(bson.M{"buildno": house.BuildNo}).Count()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	if count > 0 {
		return "编号为" + house.BuildNo + "的建筑已经存在"
	}
	err = c.Insert(&house)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindHouses 查询小区信息集
func FindHouses(db *mgo.Database, buildNO string, pageNo int, pageSize int) interface{} {
	c := db.C(HouseTabelName)
	var result []House
	var err error
	if buildNO == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"buildno": buildNO}).All(&result)
	}
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}