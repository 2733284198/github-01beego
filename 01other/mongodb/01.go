package main

import (
	"fmt"
	"time"

	//"crypto/rand"
	//"encoding/json"
	//"fmt"
	"log"
	//"time"
	"math/rand"
	//"net/http"

	//"goji.io"
	//"goji.io/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type People struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	//session, err := mgo.Dial("mongodb://root:87654321@127.0.0.1:27017/douyin")
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/user")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//c := session.DB("user").C("people")

	/*index := mgo.Index{
		Key:        []string{"isbn"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}*/

	c := session.DB("user").C("people")

	// 查找单本书
	var people People
	err = c.Find(bson.M{"age": 20}).One(&people)
	if err != nil {
		log.Println("查找isbn的书失败:", "age:20")
	} else {
		log.Println("查找isbn的书:", people)
	}

	// 修改书
	people.Name = "manlan-manlan1"
	people.Age = "21-31"

	err = c.Update(bson.M{"age": 21}, &people)
	if err != nil {
		log.Println("修改书失败:", "age:21-31")
	} else {
		log.Println("修改书成功:")
	}

	// 删除书
	err = c.Remove(bson.M{"age": 20})
	if err != nil {
		log.Println("删除书失败:", "age:20")
	}

	// 添加书
	people.Name = "manlan-manlan1"
	//people.Age = string(rune(time.Now().Second()))
	//people.Age = string(rune(rand.Intn(50)))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	people.Age = fmt.Sprintf("%d", r.Intn(50))

	err = c.Insert(people)
	if err != nil {
		log.Println("添加书失败:", "age:21-31")
	} else {
		log.Println("添加书成功:", people)
	}

	// 查找所有书籍
	var peoples []People
	err = c.Find(bson.M{}).All(&peoples)
	if err != nil {
		log.Println("查找isbn的书失败:")
	} else {
		log.Println("查找isbn的书: ", peoples)
	}
}
