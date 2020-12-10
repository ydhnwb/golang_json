package entity

//User represents user from json data
type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Profile  Profile   `json:"profile"`
	Articles []Article `json:"articles:"`
}
