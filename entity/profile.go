package entity

//Profile struct represents profile data from json
type Profile struct {
	FullName string   `json:"full_name"`
	Birthday string   `json:"birthday"`
	Phones   []string `json:"phones"`
}
