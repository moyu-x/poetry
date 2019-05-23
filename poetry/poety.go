package poetry

import (
	"encoding/json"
	"io/ioutil"
	"poetry/configs"
	"poetry/tools"
	"strings"
)

// ReadPoetryAuthors 获取诗人的信息
func ReadPoetryAuthors(dynasty string) []Author {
	poetryPath := configs.GetChinesePoetryPath()
	authorJSONPath := "/json/authors." + dynasty + ".json"
	authorJSONdata, err := ioutil.ReadFile(poetryPath + authorJSONPath)
	tools.CheckErr(err)

	authors := []Author{}

	err = json.Unmarshal([]byte(authorJSONdata), &authors)
	tools.CheckErr(err)

	return authors
}

// ReadPoetry 读取唐宋时期的诗歌
func ReadPoetry() []Poetry {
	poetryPath := configs.GetChinesePoetryPath()

	files, err := ioutil.ReadDir(poetryPath + "/json")
	tools.CheckErr(err)

	poetrys := []Poetry{}

	for _, file := range files {
		if strings.Contains(file.Name(), "poet") {
			poetryJSONPath := poetryPath + "/json/" + file.Name()

			poetryJSONData, err := ioutil.ReadFile(poetryJSONPath)
			tools.CheckErr(err)

			poetryJSONArray := []JSONPoetry{}
			err = json.Unmarshal(poetryJSONData, &poetryJSONArray)

			var popetryType string
			if strings.Contains(file.Name(), "song") {
				popetryType = "SONG"
			} else {
				popetryType = "TANG"
			}

			for _, poetryJSON := range poetryJSONArray {
				poetry := Poetry{}
				poetry.Author = poetryJSON.Author
				poetry.Title = poetryJSON.Title
				poetry.Paragraphs = strings.Join(poetryJSON.Paragraphs, "")
				poetry.Strains = strings.Join(poetryJSON.Strains, "")
				poetry.Type = popetryType
				poetrys = append(poetrys, poetry)
			}

		}
	}
	return poetrys
}
