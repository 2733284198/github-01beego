package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	//viper.AddConfigPath("./01otherlibs")
	//viper.AddConfigPath(".")

	var yaml = []byte(`
		Hacker: true
		name: steve
		hobbies:
		- skateboarding
		- snowboarding
		- go
		clothing:
		  jacket: leather
		  trousers: denim
		age: 35
		eyes : brown
		beard: true
    `)

	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	if err != nil {
		fmt.Println(err)
		return
	}
	hacker := viper.GetBool("Hacker")
	hobbies := viper.GetStringSlice("hobbies")
	jacket := viper.Get("clothing.jacket")
	age := viper.GetInt("age")
	fmt.Println("Hacker: ", hacker, ",hobbies: ", hobbies, ",jacket: ", jacket, ",age: ", age)

}
