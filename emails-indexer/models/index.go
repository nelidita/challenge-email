package models

type Index struct {
	Name         string  `json:"name"`
	Shard_num    int     `json:"shard_num"`
	Storage_type string  `json:"storage_type"`
	Mappings     Mapping `json:"mappings"`
}

type Mapping struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Id       Property `json:"id"`
	From     Property `json:"from"`
	To       Property `json:"to"`
	Content  Property `json:"content"`
	Subject  Property `json:"subject"`
	Date     Property `json:"date"`
	Filepath Property `json:"filepath"`
}

type Property struct {
	TypeProperty string `json:"type"`
	Index        bool   `json:"index"`
	Store        bool   `json:"store"`
	Sortable     bool   `json:"sortable"`
	Format       string `json:"format"`
}

type ResCreateIndex struct {
	Index        string
	Message      string
	Storage_type string
}
