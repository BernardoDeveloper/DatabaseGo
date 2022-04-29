package main

import "os"

var tableName string = "database.json"
var file, _ = os.Create(tableName)

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

	var update = []Syntax{
		{
			Id:    1,
			Key:   "name",
			Value: "Pedrinho",
		},
	}

	Create(data)
	Update(1, update)
	OneValue(1)
	Delete(2)

	All()
}
