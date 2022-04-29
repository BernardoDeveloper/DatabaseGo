package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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

func Update(id int, data []Syntax) {
	var str []Syntax
	var newData []Syntax

	dataArchive, err := ioutil.ReadFile(tableName)
	if err != nil {
		panic(err)
	}

	_ = json.Unmarshal(dataArchive, &str)

	for _, v := range str {
		if v.Id == id {
			newData = append(newData, data...)
		} else {
			newData = append(newData, v)
		}
	}

	Create(newData)
}

func Delete(id int) {
	var data []byte
	data, _ = ioutil.ReadFile(tableName)
	fmt.Printf("%s", data)

	var str []Syntax
	_ = json.Unmarshal(data, &str)

	var newData []Syntax
	for _, v := range str {
		if v.Id != id {
			newData = append(newData, v)
		}
	}

	Create(newData)
}

func DropTable() {
	_ = os.Remove(tableName)
}
