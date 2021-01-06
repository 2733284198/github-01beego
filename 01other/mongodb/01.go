package main

import (
	//"encoding/json"
	//"fmt"
	"log"
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

	var people People

	_ = c.Find(bson.M{"age": 20}).One(&people)
	log.Println("查找isbn的书失败: ", people)

}
