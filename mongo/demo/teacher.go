package demo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/qingcc/demo_tools/mongo/util"
	"log"
)

type Teacher struct {
	Username      string    `bson:"username"`
	Age           int       `bson:"age"`
	Favoritebooks []string  `bson:"favoritebooks"`
	Subject       string    `bson:"subject"`
	Students      []Student `bson:"students"`
}

type Student struct {
	Name  string   `bson:"name"`
	Id    int      `bson:"id"`
	Likes []string `bson:"likes"`
}

func InsertTeacher(t Teacher) (err error) {
	c := util.GetMgoTest("teacher")
	if err = c.Insert(t); err != nil {
		log.Printf("mongo insert failed:%s", err.Error())
	}
	return
}

func GetOneTeacher(field string, val interface{}) (t *Teacher) {
	t = &Teacher{}
	c := util.GetMgoTest("teacher")
	err := c.Find(bson.M{field: val}).One(t)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func FindAllTeacher(field string, val interface{}) (ts []Teacher) {
	ts = make([]Teacher, 0)
	c := util.GetMgoTest("teacher")
	err := c.Find(bson.M{field: val}).All(&ts)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func UpdateTeacher() {
	c := util.GetMgoTest("")
	t := Teacher{
		Username: "zhangsangai",
		Age:      25,
		Subject:  "english",
		Students: []Student{
			{
				Name:  "stu1",
				Id:    2,
				Likes: []string{"games", "music", "food", "nature"},
			},
		},
	}
	err := c.Update(bson.M{"username": "zhangsan"}, t)
	if err != nil {
		log.Printf("err:%s", err.Error())
	} else {
		log.Printf("update success")
	}
}

func DeleteTeacher() (ch int) {
	c := util.GetMgoTest("")
	changedInfo, err := c.RemoveAll(
		bson.M{
			"age": 60,
		},
	)
	if err != nil {
		log.Printf("err:%s", err.Error())
	} else {
		log.Printf("delete success")
		ch = changedInfo.Removed
	}

	return
}
