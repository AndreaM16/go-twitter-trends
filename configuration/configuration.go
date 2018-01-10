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
	Credentials struct {
		Consumer struct {
			Key    string `json:"Key"`
			Secret string `json:"Secret"`
		} `json:"Consumer"`
		Access struct {
			Token  string `json:"Token"`
			Secret string `json:"Secret"`
		} `json:"Access"`
	} `json:"Credentials"`
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