package models

type User struct {
	User_id  int    `json:"user_id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password_id int    `json:"password_id"`
	User_id     int    `json:"user_id"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	IB64        string `json:"icon_base64data"`
	Description string `json:"description"`
	Password    string `json:"password"`
}
