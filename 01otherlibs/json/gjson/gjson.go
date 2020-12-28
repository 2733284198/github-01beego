package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `{"name":
			{"first":"Janet","last":"Prichard"},
		"age":47}`

const json2 = `
		{
			"programmers": 
			[
				{
				"firstName": "Janet",
				"lastName": "McLaughlin",
				}, 
				{
				"firstName": "Elliotte",
				"lastName": "Hunter",
				}, {
				"firstName": "Jason",
				"lastName": "Harold",
				}
			]
		}
`

const json3 = `
			"programmers": 
			[
				{
				"firstName": "Janet",
				"lastName": "McLaughlin",
				}, 
				{
				"firstName": "Elliotte",
				"lastName": "Hunter",
				}, {
				"firstName": "Jason",
				"lastName": "Harold",
				}
			]
`

const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`

const complexjson = `
	{
  "widget": {
    "debug": "on",
    "window": {
      "title": "Sample Konfabulator Widget",
      "name": "main_window",
      "width": 500,
      "height": 500
    },
    "image": { 
      "src": "Images/Sun.png",
      "hOffset": 250,
      "vOffset": 250,
      "alignment": "center"
    },
    "text": {
      "data": "Click Here",
      "size": 36,
      "style": "bold",
      "vOffset": 100,
      "alignment": "center",
      "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
    }
  }
}    

`

func main() {
	simple2()
}

func simple() {
	//value := gjson.Get(json, "name.last")
	value := gjson.Get(json, "age")
	println(value.String())

	value1 := gjson.Get(jsonStream[0:1], "name")
	fmt.Println(value1)

	value2 := gjson.GetMany(jsonStream, "name")
	fmt.Println(value2)
}

func simple2() {
	result := gjson.Get(json2, "programmers.#.lastName")
	//result := gjson.Get(json3, "programmers.#.lastName")

	for _, name := range result.Array() {
		println(name.String())
	}
}

func jsonjsonget() {
	name := gjson.Get(json2, `programmers.#[lastName="Hunter"].firstName`)
	println(name.String()) // prints "Elliotte"
}
