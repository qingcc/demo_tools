package main

import (
	"github.com/qingcc/demo_tools/mongo/demo"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//add()

	//demo.MulQuery()

	//insert teacher
	tinsert()
	select {}
}

var (
	skill = []string{"Go project", "mongodb", "redis", "docker", "rpcx", "rabbitmq", "zookeeper", "shell", "git", "http", "tcp", "linux"}
)

func add() {
	go func() {
		i := 0
		for {
			item := demo.Books{
				Name:   "mongo",
				Price:  rand.Float64(),
				Author: "author_" + strconv.Itoa(int(time.Now().Unix())),
				AuthorInfo: []demo.AuthorInfo{
					{
						Username:      "zhangsan",
						Age:           rand.Intn(50),
						Phone:         "15212341" + strconv.Itoa(int(time.Now().Unix()))[:3],
						Favoritebooks: skill[:rand.Intn(11)],
					},
				},
			}
			i++
			demo.InsertBooks(item)
			if i > 10 {
				log.Println("---------end---------")
				time.Sleep(time.Hour)
			}
		}
	}()
}

func tinsert()  {
	t := demo.Teacher{
		Username:      "yang teacher",
		Age:           35,
		Favoritebooks: skill[:rand.Intn(10)],
		Subject:       "math",
		Students:      []demo.Student{
			{
				Name:"zhang san",
				Id: 003,
				Likes:[]string{"food", "games", "music", "movies", "reading"},
			},
		},
	}
	tInsert(t)
	t1 := demo.Teacher{
		Username:      "wang teacher",
		Age:           40,
		Favoritebooks: skill[:rand.Intn(10)],
		Subject:       "math",
		Students:      []demo.Student{
			{
				Name:"li si",
				Id: 006,
				Likes:[]string{"listen music", "games", "basketball", "movies", "reading"},
			},
		},
	}
	tInsert(t1)

	t2 := demo.Teacher{
		Username:      "jiang teacher",
		Age:           39,
		Favoritebooks: skill[:rand.Intn(10)],
		Subject:       "yuwen",
		Students:      []demo.Student{
			{
				Name:"wang wu",
				Id: 100,
				Likes:[]string{"sleep", "games", "basketball", "movies", "reading"},
			},
		},
	}
	tInsert(t2)

	t3 := demo.Teacher{
		Username:      "li teacher",
		Age:           48,
		Favoritebooks: skill[:rand.Intn(10)],
		Subject:       "english",
		Students:      []demo.Student{
			{
				Name:"jim",
				Id: 136,
				Likes:[]string{"english", "story", "walking", "movies", "reading"},
			},
		},
	}
	tInsert(t3)

}

func tInsert(t demo.Teacher)  {
	if err := demo.InsertTeacher(t); err != nil {
		log.Printf("insert err:%s", err.Error())
		return
	}
	log.Printf("insert success!")
	return
}

func tQuerry()  {
	t := demo.GetOneTeacher("subject", "math")
	log.Printf("query result: %+v", t)
	return
}

func tFind()  {
	t := demo.FindAllTeacher("subject", "math")
	for key, tea := range t {
		log.Printf("key:%d, result:%+v", key, tea)
	}
	return
}
