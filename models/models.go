package models

// create models
type User struct{
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Age			uint8	`json:"age"`
	Location	string	`json:"location"`
}
