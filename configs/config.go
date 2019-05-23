package configs

import (
	"io/ioutil"
	"path/filepath"
	"poetry/tools"

	"gopkg.in/yaml.v3"
)

type config struct {
	Database      databse
	ChinesePoetry chinesePoetry `yaml:"chinese-poetry"`
}

type databse struct {
	URL      string
	Port     string
	Username string
	Password string
	Db       string
}

type chinesePoetry struct {
	Path string
}

func poetryConfig() config {
	config := config{}
	file, _ := filepath.Abs("config.yml")
	yamlFile, err := ioutil.ReadFile(file)
	tools.CheckErr(err)
	err = yaml.Unmarshal([]byte(yamlFile), &config)
	tools.CheckErr(err)
	return config
}

// GetDatabaseURL 用来获取数据库连接
func GetDatabaseURL() string {
	poetryConfig := poetryConfig()
	dbURL := poetryConfig.Database.Username + ":" +
		poetryConfig.Database.Password + "@(" + poetryConfig.Database.URL +
		":" + poetryConfig.Database.Port + ")/" + poetryConfig.Database.Db +
		"?charset=utf8mb4&multiStatements=true"

	return dbURL
}

// GetChinesePoetryPath is return chines-poetry floder path
func GetChinesePoetryPath() string {
	poetryConfig := poetryConfig()
	return poetryConfig.ChinesePoetry.Path
}
