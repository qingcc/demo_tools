package demo

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/qingcc/demo_tools/mongo/util"
	"log"
)

type Books struct {
	Name       string       `bson:"name"`
	Price      float64      `bson:"price"`
	Author     string       `bson:"author"`
	AuthorInfo []AuthorInfo `bson:"authorinfo"`
}

type AuthorInfo struct {
	Username      string   `bson:"username"`
	Age           int      `bson:"age"`
	Phone         string   `bson:"phone"`
	Favoritebooks []string `bson:"favoritebooks"`
}

func InsertBooks(b Books) (err error) {
	c := util.GetMgoTest("books")
	if err = c.Insert(b); err != nil {
		log.Printf("mongo insert failed:%s", err.Error())
	}
	return
}

func QueryBooks(field string, val interface{}) (b *Books) {
	b = &Books{}
	c := util.GetMgoTest("books")
	err := c.Find(bson.M{field: val}).One(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name:", b.Name)
	fmt.Println("Price:", b.Price)
	fmt.Println("Author:", b.Author)
	return
}

func MulQuery() {
	c := util.GetMgoTest("books")
	pipe := c.Pipe([]bson.M{{"$unwind": "$authorinfo"},
		{"$match": bson.M{"authorinfo.username": "zhangsan", "age": 18}},
		{"$project": bson.M{"age": 0}}})
	resp := []bson.M{}
	err := pipe.All(&resp)
	if err != nil {
		log.Printf(err.Error())
	}
}
