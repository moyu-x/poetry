package poetry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"poetry/configs"
	"poetry/model"
	"poetry/tools"
	"strings"
)

// ReadPoetryAuthors 获取诗人的信息
func ReadPoetryAuthors(dynasty string, floderPath string) []model.Author {
	poetryPath := configs.GetChinesePoetryPath()
	var authorJSONPath string
	if "ci" == floderPath {
		authorJSONPath = "/" + floderPath + "/author." + strings.ToLower(dynasty) + ".json"
	} else {
		authorJSONPath = "/" + floderPath + "/authors." + strings.ToLower(dynasty) + ".json"
	}
	authorJSONdata, err := ioutil.ReadFile(poetryPath + authorJSONPath)
	tools.CheckErr(err)

	authors := []model.Author{}

	err = json.Unmarshal([]byte(authorJSONdata), &authors)
	tools.CheckErr(err)

	for index := range authors {
		authors[index].Type = dynasty
	}

	return authors
}

// ReadPoetry 读取唐宋时期的诗歌
func ReadPoetry(dynasty string, floderPath string) []model.Poetry {
	poetryPath := configs.GetChinesePoetryPath()

	files, err := ioutil.ReadDir(poetryPath + "/" + floderPath)
	tools.CheckErr(err)

	poetrys := []model.Poetry{}

	for _, file := range files {
		// 处理唐诗
		if strings.Contains(file.Name(), "poet") {
			poetryJSONPath := poetryPath + "/json/" + file.Name()

			poetryJSONData, err := ioutil.ReadFile(poetryJSONPath)
			tools.CheckErr(err)

			poetryJSONArray := []model.JSONPoetry{}
			err = json.Unmarshal(poetryJSONData, &poetryJSONArray)
			tools.CheckErr(err)

			// var popetryType string
			// if strings.Contains(file.Name(), "song") {
			// 	popetryType = "SONG"
			// } else {
			// 	popetryType = "TANG"
			// }

			fmt.Println(poetryJSONData)

			// for _, poetryJSON := range poetryJSONArray {
			// 	poetry := Poetry{}
			// 	poetry.Author = poetryJSON.Author
			// 	poetry.Title = poetryJSON.Title
			// 	poetry.Paragraphs = strings.Join(poetryJSON.Paragraphs, "")
			// 	poetry.Strains = strings.Join(poetryJSON.Strains, "")
			// 	poetry.Type = popetryType
			// 	poetrys = append(poetrys, poetry)
			// }

		}
	}
	return poetrys
}
