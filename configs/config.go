package configs

import (
	"io/ioutil"
	"path/filepath"
	"poetry/tools"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		URL      string
		Port     string
		Username string
		Password string
		Db       string
	}
}

func poetryConfig() Config {
	config := Config{}
	file, _ := filepath.Abs("config.yml")
	yamlFile, err := ioutil.ReadFile(file)
	tools.CheckErr(err)
	err = yaml.Unmarshal([]byte(yamlFile), &config)
	tools.CheckErr(err)
	return config
}

func GetDatabaseUrl() string {
	poetryConfig := poetryConfig()
	dbURL := poetryConfig.Database.Username + ":" +
		poetryConfig.Database.Password + "@(" + poetryConfig.Database.URL +
		":" + poetryConfig.Database.Port + ")/" + poetryConfig.Database.Db +
		"?charset=utf8mb4"

	return dbURL
}
