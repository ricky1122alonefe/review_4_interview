package main

import (
	"encoding/json"
	"log"
	"github.com/goinggo/mapstructure"
	"fmt"
	"reflect"
)

func main(){
	student := People{"jqw", 11, 18}
	data:=StructToMapDemo(student)
	log.Println(data)
}

type People struct {
	Name string `json:"name_title"`
	Age int `json:"age_size"`
	Size int `json:"size"`
}

func MapToJson(){

}

func StructToJson(){
	jsonStr := `
        {
                "name": "jqw",
                "age": 18
        }
	`
	var mapResult map[string]interface{}
	if err:=json.Unmarshal([]byte(jsonStr),&mapResult);err!=nil{
		log.Println(err.Error())
	}
	log.Println(mapResult)
}

func MapToJsonDemo1(){
	mapInstances := []map[string]interface{}{}
	instance_1 := map[string]interface{}{"name": "John", "age": 10,"size":1}
	instance_2 := map[string]interface{}{"name": "Alex", "age": 12,"size":1}
	mapInstances = append(mapInstances, instance_1, instance_2)

	jsonStr, err := json.Marshal(mapInstances)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}

func MapToStructDemo(){
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "jqw"
	mapInstance["Age"] = 18

	var people People
	err := mapstructure.Decode(mapInstance, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}

func StructToMapDemo(obj interface{}) map[string]interface{}{
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
