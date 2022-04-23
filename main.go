package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var tableName string = "database.json"
var file, _ = os.Create(tableName)

type Syntax struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func isElementExist(s []Syntax, str int) Syntax {
	for _, v := range s {
		if v.Id == str {
			var data = Syntax{
				Id:    v.Id,
				Key:   v.Key,
				Value: v.Value,
			}

			return data
		}
	}

	return Syntax{}
}

func All() {
	dataArchive, err := ioutil.ReadFile(tableName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", dataArchive)
}

func OneValue(id int) {
	var data []byte
	data, _ = ioutil.ReadFile(tableName)
	fmt.Printf("%s", data)

	var str []Syntax
	_ = json.Unmarshal(data, &str)

	fmt.Print(isElementExist(str, id))
}

func Create(data []Syntax) {
	file, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile(tableName, file, 0644)
}

func Delete(id int) {
}

func DropTable() {
	_ = os.Remove(tableName)
}

func main() {
	// fmt.Print("type database file Name: ")
	// fmt.Scan(&tableName)
	// tableName += ".json"

	defer file.Close()

	var data = []Syntax{
		{
			Id:    1,
			Key:   "name",
			Value: "John",
		},
		{
			Id:    2,
			Key:   "name",
			Value: "Jane",
		},
		{
			Id:    3,
			Key:   "name",
			Value: "Jack",
		},
	}

	Create(data)
	OneValue(1)

	// All()
}
