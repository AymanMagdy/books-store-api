package bookservice

type Book struct {
	ID int   	`json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Description string `json:"description"`
}