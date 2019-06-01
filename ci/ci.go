package ci

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"poetry/configs"
// 	"poetry/poetry"
// 	"poetry/tools"
// 	"strings"
// )

// // ReadCiAuthors 用于获取宋词词人的信息
// func ReadCiAuthors() []poetry.Author {
// 	poetryPath := configs.GetChinesePoetryPath()
// 	authorJSONPath := "/ci/author.song.json"
// 	authorJSONData, err := ioutil.ReadFile(poetryPath + authorJSONPath)
// 	tools.CheckErr(err)

// 	authors := []poetry.Author{}

// 	err = json.Unmarshal(authorJSONData, &authors)
// 	tools.CheckErr(err)

// 	for _, author := range authors {
// 		author.Type = "CI"
// 	}

// 	return authors
// }

// // ReadCi 是用来读取宋词数据并转换成数据库模型
// func ReadCi() []Ci {
// 	poetryPath := configs.GetChinesePoetryPath()
// 	files, err := ioutil.ReadDir(poetryPath + "/ci")
// 	tools.CheckErr(err)

// 	cis := []Ci{}

// 	for _, file := range files {
// 		if strings.Contains(file.Name(), "ci.song") {
// 			ciJSONPath := poetryPath + "/ci/" + file.Name()

// 			ciJSONData, err := ioutil.ReadFile(ciJSONPath)
// 			tools.CheckErr(err)

// 			ciJSONArry := []JSONCi{}
// 			err = json.Unmarshal(ciJSONData, &ciJSONArry)
// 			tools.CheckErr(err)

// 			for _, ciJSON := range ciJSONArry {
// 				ci := Ci{}
// 				ci.Paragraphs = strings.Join(ciJSON.Paragraphs, "")
// 				ci.Author = ciJSON.Author
// 				ci.Rhythmic = ciJSON.Rhythmic
// 				cis = append(cis, ci)
// 			}
// 		}
// 	}
// 	return cis
// }
