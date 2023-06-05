package models

type Record struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	//Tags       []Tag
	Created_at string `json:"created_at"`
}

// type Tag struct {
// 	Name       string `json:"name"`
// 	Created_at string `json:"created_at"`
// }
