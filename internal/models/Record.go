package models

type Record struct {
	title string `json:"title"`
	body  string `json:"body"`
	tags  []Tag
}

type Tag struct {
	name `json:"name`
}
