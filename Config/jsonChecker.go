package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	DataBasePath string
	Port         string
}

func Read() Config {
	pathToFile := "Config/config.json"
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal("Config did`t find")
	}
	byteValue, _ := ioutil.ReadAll(file)
	var config Config
	err = json.Unmarshal([]byte(byteValue), &config)
	if err != nil {
		log.Fatal("Config did`t open")
	}
	err = file.Close()
	if err != nil {
		log.Fatal("Config did`t close")
	}
	return config
}
