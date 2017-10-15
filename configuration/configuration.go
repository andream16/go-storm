package configuration

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Server struct {
		Host	string	`json:"HOST"`
		Port	int   	`json:"PORT"`
		Api		string  `json:"API"`
	} `json:"SERVER"`
	Database struct {
		DRIVERNAME	string	`json:"DRIVERNAME"`
		DBNAME    	string  `json:"DBNAME"`
		USER      	string  `json:"USER"`
		HOST       	string  `json:"HOST"`
		SSLMODE    	string  `json:"SSLMODE"`
	} `json:"DATABASE"`
}

func InitConfiguration() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println("error " + err.Error())
		os.Exit(1)
	}
	return configuration
}

func getFileName() string {
	filename := "configuration.json"
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), filename)

	return filePath
}