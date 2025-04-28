package models

type UserResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}

type UserData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
