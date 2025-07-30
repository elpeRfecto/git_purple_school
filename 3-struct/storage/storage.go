package storage

import (
	"converter/3-struct/bins"
	"encoding/json"
	"fmt"
	"os"
)

func SaveBinToJSON() ([]byte, error) {
	data, err := json.MarshalIndent(bins.Bin{}, "", " ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CreateFile() {
	file, err := os.Create("bin.json")
	if err != nil {
		fmt.Println("Cannot create file")
	}
	data, err := SaveBinToJSON()
	if err != nil {
		fmt.Println(err)
	}
	file.Write(data)
	defer file.Close()
}

func JSONReader() {
	file, err := os.Open("bin.json")
	if err != nil {
		fmt.Println("Такого файла не существует")
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&bins.BinList); err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Println("Содержимое файла: ", file)
}
