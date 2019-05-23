package analects

import (
	"strings"
	"encoding/json"
	"poetry/tools"
	"io/ioutil"
	"poetry/configs"
)

// ReadAnalects 是用来读取大学的数据使用
func ReadAnalects() []Analects {
    poetryPath := configs.GetChinesePoetryPath()
    JSONPath := "/lunyu/lunyu.json"
    analectsJSONData, err := ioutil.ReadFile(poetryPath + JSONPath)
    tools.CheckErr(err)

    analectsJSON := []JSONAnalects{}

    err = json.Unmarshal(analectsJSONData, &analectsJSON)
    tools.CheckErr(err)

    analectsData := []Analects{}
    for _, item := range analectsJSON {
        analects := Analects{}
        analects.Chapter = item.Chapter
        analects.Paragraphs = strings.Join(item.Paragraphs, "")
    }

    return analectsData
}
