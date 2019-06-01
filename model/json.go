package model

// JSONCiPoetry 是词的 JSON 数据格式
type JSONCiPoetry struct {
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Rhythmic   string   `json:"rhythmic"`
}

// JSONPoetry 是诗的 JSON 数据格式
type JSONPoetry struct {
	Strains    []string `json:"strains"`
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Titile     string   `json:"title"`
}
